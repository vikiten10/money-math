package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/vikiten10/money-math/internal/auth"
	"github.com/vikiten10/money-math/internal/database"
)

const port int16 = 3000

func main() {
	// Logger configuration
	handler := slog.NewTextHandler(os.Stdout, nil)
	logger := slog.New(handler)

	// Database configuration
	db, err := database.NewSqliteDb("money_math.db")
	if err != nil {
		logger.Error("error in creating/connecting to the database", slog.Any("error", err))
		os.Exit(1)
	}

	err = database.RunMigrations(db)
	if err != nil {
		logger.Error("error in running migrations", slog.Any("error", err))
		os.Exit(1)
	}

	// HTTP server configuration
	mux := http.NewServeMux()

	mux.Handle("/api/", http.StripPrefix("/api", auth.GetAuthRoutesMux()))

	listeningAddress := fmt.Sprintf(":%d", port)

	logger.Info("starting the server", slog.String("address", listeningAddress))

	err = http.ListenAndServe(listeningAddress, mux)
	if err != nil {
		logger.Error("error in starting the server", slog.Any("error", err))
		os.Exit(1)
	}
}
