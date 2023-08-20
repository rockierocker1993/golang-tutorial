package main

import "fmt"

func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)
	firstName1, _, _ := getFullName1()
	fmt.Println(firstName1)
}

func getFullName() (string, string) {
	return "Ikhsan", "Fadly"
}

func getFullName1() (string, string, string) {
	return "Ikhsan", "Fadly", "Ganteng"
}
