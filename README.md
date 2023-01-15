# loggit

This pkg is for easy logging of errors and information. It uses the log standard flags.

## Installation
`go get github.com/JoeAttard90/loggit`

## Quick Start
Call the `NewLoggit()` function to create a new logger. Instancing a new logger at the start of a handler will allow you
to trace the logs for each request as they have an associated txID. 
```go
func NewLoggit() *Loggit {
    return &Loggit{
        txID: uuid.NewString(),
        logger:        log.New(os.Stderr, "", log.LstdFlags),
    }
}
```

## Example
```go
logger := loggit.NewLoggit()
logger.Info("retrieving id for user %q", username)
logger.Error("could not find id for user: %q in the db", err, username)
```

```text
Output:
2023/01/14 19:03:07 [INFO]: txid:2ffb0b9a-2c19-4712-8781-b5c590dbd1a7: retrieving id for user "joe"
2023/01/14 19:03:07 [ERROR] txid:2ffb0b9a-2c19-4712-8781-b5c590dbd1a7: could not find id for user: "joe" in the db: sql: no rows in result set

```

