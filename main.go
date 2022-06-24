package main

import "fmt"

type Hex int

func (h Hex) String() string {
	return fmt.Sprintf("0x%x", int(h))
}

func main() {
	fmt.Println("hello, world")
	var h Hex
	h = 2
	fmt.Println(h.String())
}
