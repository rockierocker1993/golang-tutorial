package main

import "fmt"

func main() {
	var name string
	name = "ikhsan"
	switch name {
	case "ikhsan":
		fmt.Println("Hello " + name)
	case "joko":
		fmt.Println("Hello " + name)
	default:
		fmt.Println("Hello " + name)
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("true")
	case false:
		fmt.Println("false")
	}

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("10")
	case length > 5:
		fmt.Println("5")
	default:
		fmt.Println(length)
	}

}
