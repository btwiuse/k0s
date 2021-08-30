# safety

<a href="https://www.buymeacoffee.com/mjwhitta">🍪 Buy me a cookie</a>

[![Go Report Card](https://goreportcard.com/badge/gitlab.com/mjwhitta/safety)](https://goreportcard.com/report/gitlab.com/mjwhitta/safety)

## What is this?

This Go module provides some thread-safe utilities.

## How to install

Open a terminal and run the following:

```
$ go get --ldflags="-s -w" --trimpath -u gitlab.com/mjwhitta/safety
```

## Usage

```
package main

import (
    hl "gitlab.com/mjwhitta/hilighter"
    "gitlab.com/mjwhitta/safety"
)

func main() {
    var m *safety.Map = safety.NewMap()
    var s *safety.Set = safety.NewSet()
    var sl *safety.Slice = safety.NewSlice()

    // Maps
    m.Put("a", "asdf")
    m.Put("b", "blah")
    m.Put("s", "stop")
    m.Put("t", "test")

    m.Range(
        func(k, v interface{}) bool {
            switch v.(string) {
            case "stop":
                return true
            }

            hl.Printf("%s = %s\n", k.(string), v.(string))
            return false
        },
    )

    hl.Println(m.Keys())

    // Sets
    s.Add("asdf")
    s.Add("blah")
    s.Add("stop")
    s.Add("test")

    s.Range(
        func(entry interface{}) bool {
            switch entry.(string) {
            case "stop":
                return true
            }

            hl.Printf("Set includes: %s\n", entry.(string))
            return false
        },
    )

    hl.Println(s.Get())

    // Slices
    sl.Append("asdf")
    sl.Append("blah")
    sl.Append("stop")
    sl.Append("test")

    sl.Range(
        func(i int, v interface{}) bool {
            switch v.(string) {
            case "stop":
                return true
            }

            hl.Printf("Slice %d: %s\n", i, v.(string))
            return false
        },
    )

    hl.Println(sl.Get(0))
}
```

## Links

- [Source](https://gitlab.com/mjwhitta/safety)
