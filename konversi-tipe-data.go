package main

import "fmt"

func main() {

	var nilai32 = 3728
	var nilai64 = int64(nilai32)

	fmt.Println(nilai64)

	var nilai16 = int16(nilai64)
	fmt.Println(nilai16)

	var name string = "ikhsan"
	var c byte = name[0]
	var cString = string(c)

	fmt.Println(c)
	fmt.Println(cString)

}
