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

You can store any value associated with a string key and set time after which it will expire:

```go
c.Set("1", 20, time.Second*3)
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

## Contributing

Feel free to fork the repository and submit pull requests for any improvements or features you think would be useful.