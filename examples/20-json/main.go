// Demonstrates JSON marshaling, unmarshaling, tags, custom marshaling, and streams.
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email,omitempty"`
	Token string `json:"-"`
}

type Date struct{ time.Time }

func (d Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Format("2006-01-02"))
}

func (d *Date) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	d.Time = t
	return nil
}

func main() {
	u := User{ID: 1, Name: "Alice", Token: "secret"}
	b, err := json.Marshal(u)
	fmt.Println("marshal:", string(b), err)

	pretty, _ := json.MarshalIndent(u, "", "  ")
	fmt.Println("pretty:", string(pretty))

	// Unmarshal
	var decoded User
	err = json.Unmarshal([]byte(`{"id":2,"name":"Bob","email":"b@x.com","unknown":1}`), &decoded)
	fmt.Printf("unmarshal: %+v err=%v\n", decoded, err)

	// Disallow unknown fields
	var strict User
	dec := json.NewDecoder(strings.NewReader(`{"id":3,"bad":true}`))
	dec.DisallowUnknownFields()
	err = dec.Decode(&strict)
	fmt.Println("strict error:", err)

	// Custom marshaling
	record := struct {
		When Date `json:"when"`
	}{When: Date{Time: time.Date(2026, 11, 17, 0, 0, 0, 0, time.UTC)}}
	jb, _ := json.Marshal(record)
	fmt.Println("custom date:", string(jb))

	var back struct {
		When Date `json:"when"`
	}
	json.Unmarshal(jb, &back)
	fmt.Println("parsed date:", back.When.Format("2006-01-02"))

	// Streaming encoder (JSON Lines)
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	for _, name := range []string{"A", "B"} {
		enc.Encode(User{Name: name})
	}
	fmt.Println("stream:", buf.String())

	// map[string]any — numbers become float64
	var generic map[string]any
	json.Unmarshal([]byte(`{"n":42}`), &generic)
	fmt.Printf("generic number: %v type=%T\n", generic["n"], generic["n"])
}
