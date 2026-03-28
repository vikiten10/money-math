# Architecture

## Goals

- Single deployable binary
- Minimal backend dependencies
- SQLite as the only persistence layer
- Embedded SPA served by the backend
- Clear separation between domain logic, persistence, and delivery

## Final Technical Stack

### Backend

- Language: Go
- HTTP: standard library
- Persistence: SQLite
- DB access: `database/sql` with raw SQL
- Migrations: embedded SQL files
- Asset serving: `embed`

### Frontend

- Framework: Vue 3
- Language: TypeScript
- Build tool: Vite
- Testing: Vitest
- UI primitives: Reka UI
- Rendering model: SPA

## Why This Stack

### Go + stdlib HTTP

This keeps the backend small, portable, and easy to reason about. It also aligns with the goal of minimal dependencies.

### SQLite

SQLite fits the single-user, self-hosted beta well. It simplifies deployment, backup, and local development.

### Embedded SPA

The backend will serve the built frontend assets directly from the Go binary. This preserves the one-binary deployment target.

### Vue + TypeScript

Vue matches current familiarity and should reduce frontend delivery friction compared to switching ecosystems.

### Reka UI

Reka UI provides accessible primitives without forcing a heavy visual framework or an overly generic dashboard style.

## High-Level Backend Shape

The backend should evolve into a few clear modules:

- `internal/auth`
- `internal/accounts`
- `internal/transactions`
- `internal/plans`
- `internal/reports`
- `internal/database`
- `internal/web`

Suggested responsibilities:

- Domain modules own validation, core rules, and persistence queries
- `internal/database` owns DB setup and migrations
- `internal/web` owns embedded SPA serving and shared HTTP concerns

## API Direction

- JSON API under `/api`
- Session-based auth for beta
- Version routes only when needed

The beta should avoid building an overly abstract API layer. Keep endpoints explicit and close to the product workflows.

## Frontend Direction

The frontend should prioritize:

- Fast transaction entry
- Simple account overviews
- Clear planned-transaction flows
- Default reports with minimal configuration

Use reusable domain-facing components rather than over-abstracting a design system too early.

## Post-Beta Extension Areas

These are out of beta, but the architecture should leave room for them:

- Receipt parsing with LLMs
- External connectors such as Telegram
- Bulk ingestion and automation pipelines
- Custom report builder or data playground
