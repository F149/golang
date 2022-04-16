package main

import "fmt"

func Names() (string, string) {
	return "Foo", "Bar"
}

func Nick() (first string, second string) {
	first = "First"
	second = "Second"
	return
}

func main() {
	n1, n2 := Names()
	fmt.Println(n1, n2)

	n3, _ := Names()
	fmt.Println(n3)

	_, n4 := Names()
	fmt.Println(n4)

	n5, n6 := Nick()
	fmt.Println(n5, n6)
}
