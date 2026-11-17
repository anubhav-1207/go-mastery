# Chapter 10 — Structs

A **struct** is a collection of named fields — Go's tool for building custom data types. It is the closest thing to a "class" without inheritance.

## Declaring and using structs

```go
type Person struct {
    Name string
    Age  int
}

p := Person{Name: "Alice", Age: 30}
p.Age = 31

fmt.Println(p.Name, p)

// Positional (order-dependent — use sparingly)
q := Person{"Bob", 25}

// Zero value
var z Person   // {"", 0}
```

Nested structs:

```go
type Address struct {
    City  string
    Zip   string
}

type User struct {
    Name    string
    Address Address
}

u := User{Name: "Alice", Address: Address{City: "Paris", Zip: "75001"}}
fmt.Println(u.Address.City)
```

## Struct tags

Tags are string metadata attached to fields, read via reflection by libraries like `encoding/json`:

```go
type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email,omitempty"`
    Secret string `json:"-"`
}
```

- `omitempty` — omit the field if it has the zero value.
- `json:"-"` — never marshal this field.
- Tags are just strings; they have no effect unless a library interprets them.

## Anonymous structs

For one-off data:

```go
point := struct {
    X, Y int
}{X: 1, Y: 2}
```

Useful in tests for table-driven cases.

## Embedded structs (composition)

Embedding promotes the fields and methods of one struct into another:

```go
type Animal struct {
    Name string
}

func (a Animal) Speak() string { return "..." }

type Dog struct {
    Animal          // embedded field
    Breed string
}

d := Dog{Animal: Animal{Name: "Rex"}, Breed: "Labrador"}
fmt.Println(d.Name)      // promoted field → "Rex"
fmt.Println(d.Speak())   // promoted method → "..."
```

Embedding is how Go achieves "inheritance-like" reuse, but it is **composition**: `Dog` *has an* `Animal`, not *is an* `Animal` in the type system (though method promotion makes it behave similarly).

Field name conflicts are resolved by depth — the shallowest field wins.

## Value vs pointer semantics

Structs are **values**:

```go
a := Person{"Alice", 30}
b := a       // copy
b.Age = 40
fmt.Println(a.Age) // 30 — unchanged
```

To share, use a pointer:

```go
b := &a
b.Age = 40
fmt.Println(a.Age) // 40 — changed
```

When passing to functions:

```go
func updateValue(p Person)  { p.Age = 1 }  // copy; caller unaffected
func updatePointer(p *Person) { p.Age = 1 } // caller affected
```

**Rule:** use a pointer receiver when the method must mutate the receiver, or when the struct is large; use a value receiver for small, immutable structs (Chapter 11).

## Struct literals with pointers

```go
p := &Person{Name: "Alice"}   // pointer to a struct literal
```

Go allows `p.Name` without explicit dereference — syntactic sugar for `(*p).Name`.

## Comparison

Structs are comparable **if all fields are comparable**:

```go
type Point struct{ X, Y int }
fmt.Println(Point{1,2} == Point{1,2})  // true

type S struct{ Data []int }
fmt.Println(S{} == S{})  // ERROR: slices are not comparable
```

Use `slices.Equal` or `reflect.DeepEqual` for non-comparable fields.

## Memory layout and alignment

Fields are laid out in declaration order with padding for alignment. Ordering fields from largest to smallest often reduces padding:

```go
// Less efficient padding
type Bad struct {
    A bool    // 1 byte + 7 padding
    B int64   // 8 bytes
}

// Better
type Good struct {
    B int64   // 8 bytes
    A bool    // 1 byte + 7 trailing padding
}
```

The Go compiler and GC handle this; optimize only when profiling shows memory pressure (Chapter 24).

## Constructors

Go has no constructors, but returning a pointer from a `New...` function is idiomatic:

```go
func NewPerson(name string, age int) *Person {
    return &Person{
        Name: name,
        Age:  age,
    }
}
```

Use a constructor when you need validation or defaults; otherwise a literal is fine.

## JSON example

```go
type Product struct {
    ID    int     `json:"id"`
    Name  string  `json:"name"`
    Price float64 `json:"price"`
}

// json.Marshal(Product{1, "Widget", 9.99})
// {"id":1,"name":"Widget","price":9.99}
```

## Runnable example

See `examples/09-structs/main.go`.

## Chapter summary

- Structs group named fields into a new type: `type T struct { ... }`.
- Tags provide metadata for reflection-based libraries (notably JSON).
- Embedding promotes fields/methods — Go's composition mechanism.
- Structs are values; use pointers to mutate or avoid large copies.
- Comparable structs can be compared with `==` if all fields are comparable.

## Exercises

1. Define a `Rectangle` struct with `Width` and `Height` and a function that computes area.
2. Show that assigning a struct copies it, while assigning a pointer shares it.
3. Embed a `Vehicle` struct in a `Car` struct and access promoted fields.
4. Add JSON tags to a struct so a field is omitted when empty.

*(Solutions: `exercises/SOLUTIONS.md`, Chapter 10)*
