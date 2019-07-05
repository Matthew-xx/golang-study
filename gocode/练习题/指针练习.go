package main

import "fmt"

type Person struct {
	name    string
	marry   bool
	age     int
	classes []string
}

func initPerson0(person *Person) {
	person.name = "小明"
	person.marry = true
	person.age = 20
	person.classes = []string{"数学", "英语", "语文"}
}
func initPerson1(person **Person) {
	*person = &Person{
		"小红",
		true,
		14,
		[]string{
			"数学",
			"物理",
			"化学",
		},
	}
}
func main() {
	var person *Person
	// initPerson0(person)需要避免空指针
	initPerson1(&person)
	fmt.Println("person = ", *person)
}
