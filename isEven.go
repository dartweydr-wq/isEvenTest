package main

import "fmt"

func IsEven(num int) bool {
	return num%2 == 0
}

func main() {
	fmt.Println(IsEven(2))
}
