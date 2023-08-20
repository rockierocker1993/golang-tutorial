package main

import "fmt"

func main() {

	for i := 0; i <= 10; i++ {
		if i == 2 {
			continue
		}
		if i == 8 {
			break
		}
		fmt.Println("Perulangan ", i)
	}
}
