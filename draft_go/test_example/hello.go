package main

import "fmt"

func GetName() string {
	return "Golang best"
}

func main() {
	name := GetName()
	fmt.Println("Hello ", name)
}
