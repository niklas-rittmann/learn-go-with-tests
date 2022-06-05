package main

import "fmt"

const englischHelloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}

	return englischHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Niklas"))
}
