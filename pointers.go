package main

import "fmt"

func returnStr() string {
	return "string"
}

func returnNil() *string {
	// you cannot return nil from a function that returns a string
	// but you can return nil from a function that returns a pointer to a string
	// the zero value of a pointer is nil
	return nil
}

func returnAmp() *string {
	// &"string" doesn't work
	s := "string"
	return &s
}

func main() {
	// prints string
	fmt.Println(returnStr())
	// prints <nil>
	fmt.Println(returnNil())
	// prints a memory address
	fmt.Println(returnAmp())
	// make s hold a memory address
	s := returnAmp()
	// prints a memory address
	fmt.Println(s)
	// make sp hold the value stored in the memory address held by s
	sp := *s
	// prints string
	fmt.Println(sp)
}
