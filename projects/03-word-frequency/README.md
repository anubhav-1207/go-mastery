# Project 03 — Word Frequency

**Skills:** maps, `strings.FieldsFunc`, runes/unicode, sorting, deterministic output.

## Run

```bash
go run ./projects/03-word-frequency
go run ./projects/03-word-frequency -file README.md -top 5
```

## Test

```bash
go test ./projects/03-word-frequency
```

## What to notice

- Map iteration order is random, so `Top` sorts explicitly for deterministic results.
- Ties are broken alphabetically for stable output.
- `strings.FieldsFunc` with `unicode.IsLetter/IsDigit` handles punctuation correctly.
