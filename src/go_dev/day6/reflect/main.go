package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string
	Age   int
	Score float32
}

func (s Student) print1() {
	fmt.Println(s)
}

func test(b interface{}) {
	// 类型
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.ValueOf(b))

	v := reflect.ValueOf(b)
	// 类别
	k := v.Kind()
	fmt.Println(k)

	iv := v.Interface()
	stu, ok := iv.(Student)
	if ok {
		fmt.Printf("%s %T", stu, stu)
	}
}

func testInt(b interface{}) {
	val := reflect.ValueOf(b)
	val.Elem().SetInt(100)

	c := val.Elem().Int()
	fmt.Println(c)
}

func testStruct(a interface{}) {
	val := reflect.ValueOf(a)
	kd := val.Kind()
	if kd != reflect.Struct {
		fmt.Println("error")
		return
	}

	num := val.NumField()
	methNum := val.NumMethod()
	fmt.Println(num, methNum)

	fmt.Println(val.NumMethod())

	//var params []reflect.Value
	//val.Method(0).Call(params)
}

func main() {
	//var a int = 200
	//test(a)
	//
	//var s Student = Student{
	//	Name:  "lili",
	//	Age:   2,
	//	Score: 3,
	//}
	//
	//test(s)

	//var b int = 1
	//testInt(&b)
	//fmt.Println(b)

	var a Student = Student{
		Name:  "lili",
		Age:   0,
		Score: 1,
	}

	testStruct(a)
}
