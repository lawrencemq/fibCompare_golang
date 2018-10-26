package main

import (
	"fmt"
	"./fibs"
)

func main() {

	fmt.Println(fibs.FibLinkedList(1))
	fmt.Println(fibs.FibLinkedList(2))
	fmt.Println(fibs.FibLinkedList(3))
	fmt.Println(fibs.FibLinkedList(4))
	fmt.Println(fibs.FibLinkedList(5))
	fmt.Println(fibs.FibLinkedList(10))

	fmt.Println(fibs.FibRecursive(10))
	fmt.Println(fibs.FibTwoNumbers(10))
}
