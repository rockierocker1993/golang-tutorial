package main

import "fmt"

func main() {
	firstName, lastName, middleName := namedFunction()
	firstName2, lastName2, middleName2 := namedFunction2()
	fmt.Println(firstName, middleName, lastName)
	fmt.Println(firstName2, middleName2, lastName2)
	firstName, lastName, middleName = namedFunction()
	fmt.Println(firstName, middleName, lastName)
}

func namedFunction() (firstName, lastName, middleName string) {
	firstName = "ikhsan"
	middleName = "fadly'"
	lastName = "ganteng"
	return firstName, lastName, middleName
}
func namedFunction2() (firstName, lastName, middleName string) {
	firstName = "ikhsan"
	middleName = "fadly'"
	lastName = "ganteng"
	return
}
