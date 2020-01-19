package main

import "fmt"

// 值类型
type Student struct {
	name string
	age  int
}

// 指针类型
type Test interface {
	Print()
}

func (this Student) Print() {
	fmt.Println(this.age, this.name)
}

func main() {
	a := Student{
		name: "lili",
		age:  0,
	}
	a.Print()

	var t Test
	var stu Student = Student{
		name: "11111",
		age:  11,
	}

	t = stu
	t.Print()
}
