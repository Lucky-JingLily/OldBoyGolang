package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Carer interface {
	GetName() string
	Run()
	DiDi()
}

type BMW struct {
	Name string
}

type Reader interface {
	Read()
}

type Write interface {
	Write()
}

type ReadWrite interface {
	Reader
	Write
}

type File struct {
}

func (f *File) Read() {
	fmt.Println("read data")
}

func (b *BMW) GetName() string {
	return b.Name
}

func (b *BMW) Run() {
	fmt.Println(b.Name, "Run")
}

func (b *BMW) DiDi() {
	fmt.Println(b.Name, "DiDi")
}

type BYD struct {
	Name string
}

func (b *BYD) GetName() string {
	return b.Name
}

type Student struct {
	Name string
	Id   string
	Age  int
}

type Book struct {
	Name   string
	Author string
}

type StudentArray []Student

func (p StudentArray) Len() int {
	return len(p)
}

func (p StudentArray) Less(i, j int) bool {
	return p[i].Name > p[j].Name
}

func (p StudentArray) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	//var car Carer
	//fmt.Println(car)
	//
	//var bmw BMW = BMW{Name:"ll"}
	//car = &bmw
	//
	//fmt.Println(car.GetName())
	//car.Run()
	//car.DiDi()
	//
	//var byd BYD = BYD{Name:"byd"}
	//car = &byd

	var stus StudentArray
	for i := 0; i < 10; i++ {
		stu := Student{
			Name: fmt.Sprintf("stu%d", rand.Intn(100)),
			Id:   fmt.Sprintf("id%d", rand.Intn(100)),
			Age:  rand.Intn(20),
		}
		stus = append(stus, stu)
	}

	for _, v := range stus {
		fmt.Println(v)
	}

	sort.Sort(stus)

	fmt.Println("paixu hou -----")

	for _, v := range stus {
		fmt.Println(v)
	}

	var f *File
	var b interface{}
	b = f
	v, ok := b.(Reader)

	fmt.Println(v, ok)
}
