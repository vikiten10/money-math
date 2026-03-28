# Money Math

Money Math is an open source personal finance management web application built for a single-user, self-hosted workflow first.

The project is intentionally opinionated:

- Go backend with minimal third-party dependencies
- SQLite as the only persistence store
- Vue SPA frontend embedded into the Go binary
- One deployable binary that serves both the API and the web UI
- Beta scope optimized for personal bookkeeping instead of generic finance features

## Status

Early development. The current goal is to ship a usable beta with a stable core data model and a simple deployment story.

## Final Stack

### Backend

- Go
- Standard library HTTP server and routing
- SQLite
- Raw SQL through `database/sql`
- Embedded SQL migrations
- Embedded frontend assets
- Minimal backend dependencies

### Frontend

- Vue 3
- TypeScript
- Vite
- Vitest
- Reka UI
- SPA served by the Go backend

## Beta Scope

The first beta is intentionally narrow.

### Included

- Single-user authentication
- Account management
  - Banks
  - Credit cards
  - Loans
  - Cash
  - Other custom accounts
- Transactions and bookkeeping
- Planned transactions
  - One-time transactions
  - Recurring transactions
  - Income
  - Expense
  - Transfer
- Default reports only

### Explicitly Out Of Beta

- Multi-user support
- Bank sync and aggregation
- Receipt parsing with LLMs
- External connectors such as Telegram
- Custom report builder or data playground

These are likely post-beta features, but the architecture should leave room for them.

## Project Principles

- Keep the backend small and understandable
- Prefer stdlib and simple code over framework-heavy abstractions
- Preserve a clean path to a single deployable binary
- Build for real personal usage first
- Keep the first beta focused and stable

## Planned Documentation

- [Architecture](docs/architecture.md)
- [Beta Scope](docs/beta-scope.md)

## Development Direction

The short-term implementation order is:

1. Single binary web server with embedded SPA
2. Single-user auth
3. Accounts
4. Transactions
5. Planned transactions
6. Default reports

## License

[MIT](LICENSE)
