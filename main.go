package main

import "fmt"

func main() {

	Sprintf()

}

// Sprintf example
func Sprintf() {

	s := fmt.Sprintf("http://%s:1234/path", "127.0.0.1")

	fmt.Println("formatted string:", s)

}
