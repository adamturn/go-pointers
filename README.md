# Go Pointers
In case I forget: https://tour.golang.org/moretypes/1

## `&`
The `&` operator generates a pointer to its operand.
```golang
// i holds value
var i int
i = 42
// p holds memory address
var p *int
p = &i
```
## `*`
The `*` operand denotes the value stored at the memory address held by the pointer.
```golang
*p = 420
if i == 420 {
    fmt.Println("Is this the meaning of life?")
}
```
## Pointers to structs
"Struct fields can be accessed through a struct pointer. To access a specific field of struct pointer p, we could write (*p).X. However, that notation is cumberson, so the language permits us to write just p.X, without the explicit dereference." - A Tour of Go
