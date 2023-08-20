package main

import "fmt"

func main() {
	hello := getHello("ikhsan")
	fmt.Println(hello)
}

func getHello(name string) string {
	return "Hello " + name
}
