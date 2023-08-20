package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "ikhsan"
	names[1] = "fadly"
	fmt.Println(names)
	fmt.Println(names[0])
	var angka = [...]int16{
		1, 2, 3,
	}
	fmt.Println(angka)
	fmt.Println("length", len(names))
	fmt.Println("kapasitas", cap(names))

}
