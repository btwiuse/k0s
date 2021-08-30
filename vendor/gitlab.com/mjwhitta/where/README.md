# where

<a href="https://www.buymeacoffee.com/mjwhitta">🍪 Buy me a cookie</a>

[![Go Report Card](https://goreportcard.com/badge/gitlab.com/mjwhitta/where)](https://goreportcard.com/report/gitlab.com/mjwhitta/where)

## What is this?

A `which` or `command -v` like tool for Go.

## How to install

Open a terminal and run the following:

```
$ go get --ldflags "-s -w" --trimpath -u gitlab.com/mjwhitta/where
$ go install --ldflags "-s -w" --trimpath \
    gitlab.com/mjwhitta/where/cmd/where@latest
```

Or install from source:

```
$ git clone https://gitlab.com/mjwhitta/where.git
$ cd where
$ git submodule update --init
$ make install
```

**Note:** `make install` will install to `$HOME/.local/bin`.

## Usage

```
package main

import (
    "fmt"

    "gitlab.com/mjwhitta/where"
)

func main() {
    fmt.Println(where.Is("bash"))
}
```

## Links

- [Source](https://gitlab.com/mjwhitta/where)
