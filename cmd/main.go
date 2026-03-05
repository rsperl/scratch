package main

import "fmt"

var password = "super-secret"

var highEntropyString = "aHR0cHM6Ly9leGFtcGxlLmNvbS9wYXRoP3F1ZXJ5PWNyeXB0b2dyYXBoeSZyZXN1bHQ9c2VjcmV0"

func main() {
	fmt.Println("Hello, World!")
	fmt.Printf("Your password is %s\n", password)
	fmt.Printf("High entropy string: %s\n", highEntropyString)
}
