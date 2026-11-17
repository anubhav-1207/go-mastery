# Chapter 21 — JSON

The `encoding/json` package marshals Go values to JSON and unmarshals back. It uses struct tags to map fields.

## Marshaling

```go
type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email,omitempty"`
}

u := User{ID: 1, Name: "Alice"}

b, err := json.Marshal(u)
// {"id":1,"name":"Alice"}

b, err := json.MarshalIndent(u, "", "  ")  // pretty-printed
```

`omitempty` drops zero values (`0`, `""`, `nil`, `false`, empty slice/map).

## Unmarshaling

```go
var u User
err := json.Unmarshal([]byte(`{"id":2,"name":"Bob"}`), &u)
```

You **must** pass a pointer — otherwise the function cannot modify your variable.

Unknown fields are ignored by default. Use a decoder to reject them:

```go
dec := json.NewDecoder(reader)
dec.DisallowUnknownFields()
err := dec.Decode(&u)
```

## Struct tags in detail

```go
type Record struct {
    FieldA string `json:"fieldA"`          // rename
    FieldB string `json:"field_b"`         // snake_case output
    Hidden string `json:"-"`               // never encode/decode
    Opt    string `json:"opt,omitempty"`   // omit when empty
    Str    int    `json:"str,string"`      // encode int as JSON string
}
```

Unexported (lowercase) fields are always ignored by `encoding/json`.

## `any` and nested JSON

```go
var data map[string]any
json.Unmarshal(payload, &data)
```

JSON numbers decode as `float64` by default. To preserve integer precision, use `json.Decoder.UseNumber()`:

```go
dec.UseNumber()
```

## Custom marshaling

Implement `MarshalJSON` / `UnmarshalJSON` for full control:

```go
type Date struct{ time.Time }

func (d Date) MarshalJSON() ([]byte, error) {
    return json.Marshal(d.Format("2006-01-02"))
}

func (d *Date) UnmarshalJSON(b []byte) error {
    var s string
    if err := json.Unmarshal(b, &s); err != nil { return err }
    t, err := time.Parse("2006-01-02", s)
    if err != nil { return err }
    d.Time = t
    return nil
}
```

Or use struct fields with `time.Time` — it already marshals to RFC 3339.

## Streaming encoders/decoders

For large data or HTTP bodies, stream instead of loading everything into memory:

```go
// Encode a stream of values
enc := json.NewEncoder(os.Stdout)
for _, u := range users {
    if err := enc.Encode(u); err != nil { return err }
}

// Decode a stream of values
dec := json.NewDecoder(reader)
for {
    var u User
    if err := dec.Decode(&u); err == io.EOF {
        break
    } else if err != nil {
        return err
    }
    process(u)
}
```

`json.NewEncoder` writes a newline after each value — convenient for JSON Lines (NDJSON).

## HTML escaping

`json.Marshal` escapes `<`, `>`, `&` by default (to prevent XSS in HTML contexts). Disable with an encoder:

```go
enc := json.NewEncoder(w)
enc.SetEscapeHTML(false)
```

## Common pitfalls

1. Forgetting `&` in `Unmarshal` → no error, but no data.
2. Relying on field order — JSON objects are unordered.
3. Using `map[string]any` for structured data → loses type safety.
4. Expecting `json.Number` without `UseNumber()` → you get `float64`.
5. Marshaling channels/functions → `unsupported type` error.

## Performance notes

`encoding/json` uses reflection and is not the fastest. For performance-critical paths consider `encoding/json/v2` (newer Go), `github.com/goccy/go-json`, or code generation (`easyjson`, `ffjson`). Measure before optimizing.

## Runnable example

See `examples/20-json/main.go`.

## Chapter summary

- `json.Marshal`/`Unmarshal` convert between Go values and JSON bytes.
- Struct tags control names, omission, and special encodings.
- Always pass a pointer to `Unmarshal`.
- Use `json.Encoder`/`Decoder` for streaming and HTTP bodies.
- Prefer typed structs over `map[string]any`.

## Exercises

1. Marshal a struct with `omitempty` tags and inspect the output.
2. Unmarshal JSON into a struct and handle unknown fields with `DisallowUnknownFields`.
3. Implement `MarshalJSON` for a custom `Money` type.
4. Write a stream decoder that processes newline-delimited JSON.

*(Solutions: `exercises/SOLUTIONS.md`, Chapter 21)*
