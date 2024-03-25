package main

import (
	"GomoduleTest/adder"
	"fmt"
)

func main() {
	sum := adder.Add(5, 3)
	fmt.Println(sum)
}
