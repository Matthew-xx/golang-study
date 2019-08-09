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
	//子接口可以转换为母接口，反之不行
	var person Personer
	//person:=Personer{}
	person = &Student{"mike", 666}
	var human Humaner
	//person = human //err
	human = person //可以,多的可以给少的赋值，少的不能给多的赋值
	human.sayhi()
}
