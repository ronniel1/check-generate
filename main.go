package main

import "fmt"

//go:generate ./command.sh
//go:generate echo hello
func main() {
	fmt.Println("hello world")
}
