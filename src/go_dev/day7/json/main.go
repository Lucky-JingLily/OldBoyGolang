package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score string
}

func testStruct() {
	stu := &Student{
		Name:  "lili",
		Age:   11,
		Score: "11.11",
	}
	data, err := json.Marshal(stu)
	if err != nil {
		fmt.Printf("json.marshal failed, err:", err)
		return
	}

	fmt.Println(string(data))

	var stu1 Student
	// 第二个参数记得传地址
	json.Unmarshal(data, &stu1)
	fmt.Println(stu1)
}

func main() {
	testStruct()
}
