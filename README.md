# Metastore

Metastore is a simple Go library which lets you attach a Ruby-style Hash object to any struct.

Set() stores a value, Get() attempts to retrieve it and returns a boolean indicating its success in that endeavour, and Fetch() will panic() if the requested key does not exist.

I find this to be useful for attaching arbitrary debugging data to various objects quickly.

## Installation

```
go get github.com/darktriad/metastore
```

## Usage

```go
type Foo struct {
	Metadata metastore.Metastore
}

f := Foo{}
f.Metadata.Set("some_float", 1.23)
f.Metadata.Set("some_bool", true)

// this will panic() because the key is unknown
_ = f.Metadata.Fetch("some_key")

fmt.Printf(
	"some_float: %.2f - some_bool: %t\n",
	f.Metadata.Fetch("some_float").(float64),
	f.Metadata.Fetch("some_bool").(bool),
)

// fetching an object using Get()

object, ok := f.Metadata.Get("some_key")
if !ok {
	fmt.Println("couldn't find value")
} else {
	// cast object to its real type and do something with it...
}
```
