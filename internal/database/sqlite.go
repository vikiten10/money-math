package database

import (
	"database/sql"
	"embed"
	"fmt"
	"path"
	"slices"
	"strings"

	_ "modernc.org/sqlite"
)

const migrationsFolder string = "migrations"

//go:embed migrations/*.sql
var migrations embed.FS

func NewSqliteDb(path string) (*sql.DB, error) {
	dsn := fmt.Sprintf("file:%s", path)

	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)

	return db, nil
}

func RunMigrations(db *sql.DB) error {
	if err := ensureMigrationsTableExists(db); err != nil {
		return err
	}

	entries, err := migrations.ReadDir(migrationsFolder)
	if err != nil {
		return err
	}

	fileNamesPlaceholderBuilder := strings.Builder{}
	fileNames := make([]any, 0, 256) // .Exec method requires variadic "any"
	for _, entry := range entries {
		fileNamesPlaceholderBuilder.WriteString("\nSELECT ? UNION ALL")
		fileNames = append(fileNames, entry.Name())
	}

	fileNamesPlaceholder := strings.TrimSuffix(fileNamesPlaceholderBuilder.String(), "UNION ALL")
	getUnappliedMigrationsQuery := fmt.Sprintf(
		`WITH available_migrations(file_name) AS (%s)
		SELECT am.file_name
		FROM available_migrations am
		LEFT JOIN migrations_history mh ON am.file_name = mh.file_name
		WHERE mh.file_name IS NULL;`,
		fileNamesPlaceholder)

	rows, err := db.Query(getUnappliedMigrationsQuery, fileNames...)
	if err != nil {
		return err
	}

	var unappliedMigrations []string = make([]string, 256)
	for rows.Next() {
		var fileName string

		if err = rows.Scan(&fileName); err != nil {
			return err
		}

		unappliedMigrations = append(unappliedMigrations, fileName)
	}

	if err = rows.Err(); err != nil {
		return err
	}

	insertMigrationHistoryQuery := "INSERT INTO migrations_history(file_name) VALUES (?)"

	for _, entry := range entries {
		fileName := entry.Name()
		isMigrationToBeApplied := slices.Contains(unappliedMigrations, fileName)

		if !isMigrationToBeApplied {
			continue
		}

		filePath := path.Join(migrationsFolder, fileName)
		fileContent, err := migrations.ReadFile(filePath)
		if err != nil {
			return err
		}

		tx, err := db.Begin()
		if err != nil {
			return err
		}

		defer tx.Rollback()

		_, err = tx.Exec(string(fileContent))
		if err != nil {
			return err
		}

		_, err = tx.Exec(insertMigrationHistoryQuery, fileName)
		if err != nil {
			return err
		}

		if err = tx.Commit(); err != nil {
			return err
		}
	}

	return nil
}

func ensureMigrationsTableExists(db *sql.DB) error {
	const createMigrationHistoryTableQuery string = `
	CREATE TABLE IF NOT EXISTS migrations_history (
		file_name TEXT NOT NULL PRIMARY KEY,
		applied_at TEXT NOT NULL DEFAULT (datetime('now')),
		created_at TEXT NOT NULL DEFAULT (datetime('now')),
		updated_at TEXT NOT NULL DEFAULT (datetime('now'))
	) STRICT;
	`

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		_ = tx.Rollback()
	}()

	_, err = tx.Exec(createMigrationHistoryTableQuery)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
