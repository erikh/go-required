## Check golang structs for required members

Required takes a struct and identifies any struct members with the required struct tag (the value can be anything, but must exist for it to be a valid struct tag). If the value is the zero value, it will yield an error. Otherwise it will return nil.

## Usage

```go
import (
    "github.com/erikh/go-required"
)

type MyStruct struct {
    Member string `required:""`
    OptionalMember string
}

func main() {
    m := MyStruct{}

    err := required.Required(m);
    fmt.Println(err)

    m.Member = "filled in"
    if required.Required(m) == nil {
        fmt.Println("success")
    }
}
```

## Author

Erik Hollensbe <erik@hollensbe.org>

## License

BSD 3 Clause
