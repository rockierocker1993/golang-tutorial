package main

import "fmt"

func main() {
	person := map[string]string{
		"firstName": "ikhsan",
		"lastName":  "fadly",
	}
	person["title"] = "programmer"
	fmt.Println(person)
	fmt.Println(len(person))
	delete(person, "firstName")
	fmt.Println(person)

	var book map[string]string = make(map[string]string)
	book["pemogramman"] = "golang"
	fmt.Println(book)
}
