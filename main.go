package main

import "fmt"

//go:generate echo "hello" > file.txt
//go:generate echo hello
func main() {
	fmt.Println("hello world")
}
