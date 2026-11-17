# Project 02 — Todo CLI

**Skills:** structs, methods with pointer receivers, JSON persistence, error handling, `os.Args`, `t.TempDir()`.

## Run

```bash
go run ./projects/02-todo-cli add "learn Go"
go run ./projects/02-todo-cli list
go run ./projects/02-todo-cli done 0
go run ./projects/02-todo-cli rm 0
```

Tasks are stored in `~/.go-mastery-todo.json`.

## Test

```bash
go test ./projects/02-todo-cli -v
```

## What to notice

- `Store` owns its state and exposes methods — no globals.
- `List` returns a copy so callers cannot mutate internal state.
- `NewStore` uses `errors.Is(err, os.ErrNotExist)` to distinguish a fresh start from a real error.
- Tests use `t.TempDir()` for isolated file storage.
