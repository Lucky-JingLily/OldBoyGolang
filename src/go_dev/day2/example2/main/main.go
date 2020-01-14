package main

import (
	"fmt"
	"go_dev/day2/example2/add"
	"os"
)

func main() {
	fmt.Println(add.Age)
	fmt.Println(add.Name)

	fmt.Println(fmt.Sprintf("%o", 11))

	fmt.Println(os.Environ())
	fmt.Println(os.Getenv("GOPATH"))
}
