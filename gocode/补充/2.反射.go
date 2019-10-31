package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	name string
}
type Person interface {
	Speak()
}

func (s Student) Speak() {
	fmt.Println(s.name)
}
func main() {
	// fmt.Println(nil != nil)
	// golang里面的变量分为type和value两部分
	// 接口类型的变量总是包含 value和concrete type形式

	//1. reflect.ValueOf()将接口类型转为reflect.Value类型
	//2. reflect.TypeOf()获取接口变量的类型信息，并保存至reflect.Type类型的实例中
	//3. 通过Value可以获取Type.   即reflect.TypeOf(X)==reflect.ValueOf(X).Type()
	//4. Value和Type实例，都有Kind(),返回值的底层类型 v.Kind()  t.Kind()，但也仅限底层基本类型如Struct,ptr，而
	//5. Value实例通过类型名，可以取得对应底层类型的值（仅适用于数值） v.Int(),但也仅限最大字节数，如Int64、Float64
	//6. 可以将反射实例(Value实例)，转回为接口实例。v.Interface()

	var p Person = Student{name: "张三"}
	p.Speak()
	fmt.Println(reflect.ValueOf(p).Type())
	fmt.Println(reflect.TypeOf(p))
	fmt.Println(reflect.TypeOf(p).Kind())
	fmt.Println(reflect.ValueOf(p).NumField())
	v := reflect.ValueOf(p)
	// i := v.Interface().(Person)
	i := v.Interface().(Student)
	// 7.interface()方法返回的接口实例，一般需要断言出一个concrete实际类型才可使用，但在需要传入interface的方法中，无需断言。比如fmt.Print()
	i.Speak()

	// 8. 反射实例的可设置性settable.反射实例实际存储的是原类型的副本，默认是不可设置的，若要修改，需存储指针，并获取引用
	fmt.Println("Value是否可修改", reflect.ValueOf(p).CanSet())         //false，值的副本
	fmt.Println("Value是否可修改", reflect.ValueOf(&p).CanSet())        //false,指针本身也是副本
	fmt.Println("Value是否可修改", reflect.ValueOf(&p).Elem().CanSet()) //true，指针指向的内容是同一份
	reflect.ValueOf(&p).Elem().Set(reflect.ValueOf(Student{name: "李四"}))
	p.Speak()

}
