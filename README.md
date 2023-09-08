# Cache Package

The `cache` package provides a simple in-memory key-value store for Go applications. With it, you can set, get, and delete key-value pairs easily.

## Installation

Assuming you have Go installed and set up:

```bash
go get -u github.com/exortme1ster/cache
```

## Usage

### Initialization

To start, initialize a new cache:

```go
c := cache.New()
```

### Setting a Value

You can store any value associated with a string key:

```go
c.Set("myKey", "myValue")
c.Set("anotherKey", 12345)
```

### Getting a Value

Retrieve a value by its key:

```go
value := c.Get("myKey")
if value != nil {
    fmt.Println(value) // Outputs: myValue
}
```

### Deleting a Value

Remove a key and its associated value:

```go
c.Delete("myKey")
```

## Note

This package uses a map to store data in-memory. This means:

- It's fast and suitable for temporary storage.
- If the application is stopped, all data in the cache will be lost.
- It's not thread-safe by default. If you intend to use it in a multi-threaded environment, you will need to manage concurrency control.

## Contributing

Feel free to fork the repository and submit pull requests for any improvements or features you think would be useful.