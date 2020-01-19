package main

import "fmt"

type Car struct {
	weight int
	name   string
}

type Bike struct {
	Car
	luzi int
}

type Train struct {
	Car
	lunzi int
}

func (this *Car) Run() {
	fmt.Println(this, "running")
}

func (this *Train) String() string {
	return fmt.Sprintf("name=%slunzi=%dweight%d", this.name, this.lunzi, this.weight)
}

func main() {
	var a Bike
	a.weight = 1
	a.name = "bike"
	a.luzi = 2

	var b Train
	b.weight = 2
	b.name = "train"
	b.lunzi = 22

	fmt.Println(a)
	a.Run()

	fmt.Println(b.String())
	b.Run()
	fmt.Printf("%s", &b)
}
