package main

import "fmt"

func main() {
	var angka = [...]int16{
		1, 2, 3, 4, 5, 6, 7, 8, 9,
	}
	fmt.Println(angka)
	angkaSlice1 := angka[4:]
	angkaSlice2 := angka[:5]
	angkaSlice3 := angka[2:5]
	angkaSlice4 := angka[:]

	fmt.Println(angkaSlice1)
	fmt.Println(angkaSlice2)
	fmt.Println(angkaSlice3)
	fmt.Println(angkaSlice4)
	fmt.Println(len(angkaSlice3))
	fmt.Println(cap(angkaSlice3))
	angka[3] = 11
	fmt.Println(angkaSlice3)
	appendSlice := append(angkaSlice1, 11)
	fmt.Println(appendSlice)
}
