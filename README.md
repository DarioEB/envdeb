# envdbar

A lightweight Go package for loading environment variables from .env files.

## Installation

```bash
go get github.com/db-ar/envdbar
```

## Usage

### Loading environment variables

```go
package main

import "github.com/db-ar/envdbar"

func main() {
    // Load from .env (default)
    err := envdbar.Load()
    if err != nil {
        panic(err)
    }

    // Or load from a custom file
    err = envdbar.Load(".env.production")
    if err != nil {
        panic(err)
    }
}
```

### Getting environment variables

```go
// Get a variable (returns empty string if not set)
port := envdbar.Get("PORT")

// Get a variable with a default value
port := envdbar.Get("PORT", "8080")
```

## Supported .env syntax

```env
# Comments are supported
PORT=3000
SERVER=0.0.0.0

# Quoted values (spaces preserved)
APP_NAME="My Application"
GREETING='Hello World'

# Values with equals sign
TOKEN=abc123=xyz

# Inline comments
DEBUG=true # this is a comment

# Indented lines are supported
  INDENTED_VAR=value
```

## License

MIT
