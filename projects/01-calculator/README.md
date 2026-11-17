# Project 01 — CLI Calculator

**Skills:** functions, multiple returns, errors, `switch`, `flag`, table-driven tests.

## Run

```bash
go run ./projects/01-calculator -a 10 -b 3 -op mul
# 10 mul 3 = 30
```

## Test

```bash
go test ./projects/01-calculator
```

## What to notice

- `Divide` returns `(float64, error)` instead of panicking.
- `Calculate` dispatches with a `switch` — an interface could replace it for extensibility.
- Tests are table-driven with `errors.Is` for sentinel-error checking.
