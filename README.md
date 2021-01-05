# Go Pointers
In case I forget https://tour.golang.org/moretypes/1

## `&` operates on values
The `&` operator generates a pointer to its operand.
```golang
// i holds int value
var i int
i = 42
// p holds memory address for int value
var p *int
p = &i
```
## `*` operates on pointers
The `*` operator acts like a portal to the value stored at the memory address held by its operand.
```golang
*p = 420
if i == 420 {
    fmt.Println("Is this the meaning of life?")
}
```
## Pointers to structs
Struct fields can be accessed through a struct pointer. To access a specific field of struct pointer p, we could write (*p).X. However, that notation is cumberson, so the language permits us to write just p.X, without the explicit dereference.
```golang
type vertex struct {
    X, Y int
}

func (v *vertex) mutateSelf(x, y int) {
    v.X = x
    v.Y = y
}
```
