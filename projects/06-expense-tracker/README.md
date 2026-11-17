# Project 06 — Expense Tracker

**Skills:** full CRUD app, validation, named types, JSON persistence, sorting for deterministic output, tests.

## Run

```bash
go run ./projects/06-expense-tracker add 12.50 food "lunch"
go run ./projects/06-expense-tracker add 40 transport "taxi"
go run ./projects/06-expense-tracker list
go run ./projects/06-expense-tracker summary
go run ./projects/06-expense-tracker rm 0
```

Data is stored in `~/.go-mastery-expenses.json`.

## Test

```bash
go test -race ./projects/06-expense-tracker
```

## What to notice

- `Add` validates inputs before mutating state.
- `All` returns a copy, protecting encapsulation.
- Category totals come from a map, but output is sorted for deterministic display.
- Errors are wrapped with context and compared with `errors.Is`.
- This project combines nearly every skill from chapters 01–22.
