# Chapter 20 — File I/O

Go's I/O model is built around two tiny interfaces:

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}
```

Because so many things implement them, code composes beautifully: files, network connections, buffers, compressors.

## Reading files

### Whole file (small files)

```go
data, err := os.ReadFile("notes.txt")
if err != nil {
    return err
}
fmt.Println(string(data))
```

### Opening and reading in chunks

```go
f, err := os.Open("notes.txt")
if err != nil {
    return err
}
defer f.Close()

buf := make([]byte, 4096)
for {
    n, err := f.Read(buf)
    if n > 0 {
        process(buf[:n])
    }
    if err == io.EOF {
        break
    }
    if err != nil {
        return err
    }
}
```

Always handle `n > 0` **before** checking `err` — a read may return bytes and `io.EOF` together.

### Line by line with `bufio.Scanner`

```go
f, _ := os.Open("notes.txt")
defer f.Close()

scanner := bufio.NewScanner(f)
for scanner.Scan() {
    fmt.Println(scanner.Text())
}
if err := scanner.Err(); err != nil {
    return err
}
```

Scanner has a default 64 KB token limit. Increase it for long lines:

```go
scanner.Buffer(make([]byte, 64*1024), 1024*1024)
```

For even more control, use `bufio.Reader` with `ReadString('\n')`.

## Writing files

```go
err := os.WriteFile("out.txt", []byte("hello\n"), 0644)
```

### With `os.Create` and buffered writing

```go
f, err := os.Create("out.txt")
if err != nil { return err }
defer f.Close()

w := bufio.NewWriter(f)
for i := 0; i < 1000; i++ {
    fmt.Fprintln(w, i)
}
if err := w.Flush(); err != nil {   // don't forget!
    return err
}
```

`bufio.Writer` reduces system calls dramatically. **Always flush** (or defer it) — buffered data is lost otherwise.

## File modes and permissions

```go
os.OpenFile("log.txt",
    os.O_APPEND|os.O_CREATE|os.O_WRONLY,
    0644)
```

| Flag | Meaning |
|------|---------|
| `O_RDONLY` | read only |
| `O_WRONLY` | write only |
| `O_RDWR` | read/write |
| `O_CREATE` | create if missing |
| `O_TRUNC` | truncate on open |
| `O_APPEND` | append to end |
| `O_EXCL` | fail if exists (with `O_CREATE`) |

Permissions (`0644`): owner read/write, group read, others read.

## File info and directories

```go
info, err := os.Stat("file.txt")
info.Name()
info.Size()
info.IsDir()
info.ModTime()

entries, _ := os.ReadDir(".")
for _, e := range entries {
    fmt.Println(e.Name(), e.IsDir())
}

os.Mkdir("data", 0755)
os.MkdirAll("a/b/c", 0755)
os.Remove("file.txt")
os.RemoveAll("dir")
```

## Paths

Use `path/filepath` (OS-aware) instead of string concatenation:

```go
p := filepath.Join("data", "sub", "file.txt")
filepath.Ext("a.txt")     // ".txt"
filepath.Base("/a/b.txt") // "b.txt"
filepath.Abs(".")
```

## `io` helpers

```go
io.Copy(dst, src)              // stream copy
io.ReadAll(r)                  // read everything
io.WriteString(w, "text")      // write a string
```

`io.Copy` is efficient (uses `ReaderFrom`/`WriterTo` optimizations when available).

## Temporary files

```go
f, err := os.CreateTemp("", "prefix-*.txt")
defer os.Remove(f.Name())
```

## Embedding files in the binary

```go
import _ "embed"

//go:embed templates/*.html
var templates embed.FS
```

This bakes files into the binary — great for web templates and static assets.

## Errors to expect

- `os.ErrNotExist` — file missing (check with `errors.Is`).
- `os.ErrPermission` — permission denied.
- `io.EOF` — end of input (not an error in most contexts).

## Runnable example

See `examples/19-file-io/main.go`.

## Chapter summary

- `io.Reader`/`io.Writer` are the universal I/O abstractions.
- Use `os.ReadFile`/`os.WriteFile` for small whole-file operations.
- Use `bufio.Scanner` for lines, `bufio.Writer` for efficient output (flush it!).
- Handle `n > 0` before `io.EOF`.
- Use `filepath` for portable paths; `io.Copy` for streaming.

## Exercises

1. Write a program that counts lines in a file using `bufio.Scanner`.
2. Copy a binary file with `io.Copy`.
3. Write 10,000 lines with and without `bufio.Writer` and compare timing.
4. Use `os.Stat` to print a file's size and modification time.

*(Solutions: `exercises/SOLUTIONS.md`, Chapter 20)*
