package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	url := "golang.org:80"
	connect, _ := net.Dial("tcp", url)
	fmt.Fprintf(connect, "GET / HTTP/1.0\r\n\r\n")
	status, _ := bufio.NewReader(connect).ReadString('\n')
	fmt.Println(status)
}
