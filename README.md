# go-singleton

## Get started
```go
    import "github.com/HashLoad/go-singleton"

    ...

    // Set value, value is s generic interface{}
    pattern.Singleton.Set("key", "value")

    // Get value, value is a generic interface{}
    value, err := pattern.Singleton.Get("key")
```