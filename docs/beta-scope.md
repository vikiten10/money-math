# Beta Scope

## Product Summary

Money Math beta is a self-hosted personal finance application for a single user. The beta focuses on dependable bookkeeping and reporting, not on integrations or automation-heavy workflows.

## Included Features

### Authentication

- Single-user auth
- Local session-based access

### Account Management

- Bank accounts
- Credit cards
- Loan accounts
- Cash accounts
- Other custom account types

### Transactions

- Manual transaction entry
- Income
- Expense
- Transfer
- Notes and bookkeeping metadata

### Planned Transactions

- One-time planned transactions
- Recurring planned transactions
- Planned income
- Planned expenses
- Planned transfers

### Reports

Default reports only. The first beta should ship a fixed set of useful views instead of a custom analytics builder.

Suggested default reports:

- Monthly cashflow summary
- Income vs expense trend
- Spending by category
- Spending by account
- Upcoming planned transactions

## Explicitly Out Of Scope

- Multi-user support
- Custom report builder
- Data playground
- Receipt parsing with LLMs
- Telegram integration
- Other external connectors
- Bank aggregation

## Scope Guardrails

The beta should optimize for:

- Stable bookkeeping flows
- Clear deployment
- Small backend surface area
- A maintainable domain model

The beta should avoid:

- Premature automation
- Feature sprawl
- Generic “finance platform” scope

## Post-Beta Candidate Features

- Bulk transaction entry improvements
- Receipt parsing review flows
- Telegram-based transaction capture
- Connector framework
- Customizable report views
