package main

import "fmt"

type Student struct {
	Name string
	Sex  string
}

func Test(a interface{}) {
	b, ok := a.(Student)
	if ok == false {
		fmt.Println("convert failed")
		return
	}
	//b += 3
	fmt.Println(b)
}

func just(items ...interface{}) {
	for index, v := range items {
		switch v.(type) {
		case bool:
			fmt.Printf("param %d is a bool\n", index)
		case float32, float64:
			fmt.Printf("param %d is a float\n", index)
		case int:
			fmt.Printf("param %d is a int\n", index)
		case string:
			fmt.Printf("param %d is a string\n", index)
		case Student:
			fmt.Printf("param %d is a Student", index)
		}
	}
}

func main() {
	//var a interface{}
	//var b Student
	//
	//a = b
	//
	//c := a.(int)
	//fmt.Printf("%d %T\n", c, c)
	//fmt.Printf("%d %T\n", a, a)

	var b Student
	Test(b)

	just(1, 2, 3, "string", 1.1, b)

}
