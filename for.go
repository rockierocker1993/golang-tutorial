package main

import "fmt"

func main() {
	counter := 1
	counter2 := 1
	for counter <= 10 && counter2 <= 10 {
		fmt.Println("Perulangan ke", counter)
		counter++
		counter2++
	}
	for counter3 := 1; counter3 <= 10; counter3++ {
		fmt.Println("Perulangan ke", counter3)
	}
}
