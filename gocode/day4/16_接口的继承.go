package main

import "fmt"

type Humaner interface {
	sayhi()
}
type Personer interface {
	Humaner
	sing(Irc string)
}
type Student struct {
	name string
	id   int
}

func (s *Student) sayhi() {
	fmt.Printf("student say hi:%s,%d\n", s.name, s.id)
}
func (s *Student) sing(Irc string) {
	fmt.Println("Student sing: ", Irc)
}

func main() {
	var i Personer
	s := &Student{"mike", 666}
	i = s
	i.sayhi()
	i.sing("lalala")
}
