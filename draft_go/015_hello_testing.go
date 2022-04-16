package main

import "fmt"

func GetName() string {
	return "Golang"
}

func main() {
	name := GetName()
	fmt.Println("Hello ", name)
}
