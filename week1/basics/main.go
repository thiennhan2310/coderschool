package main

import "fmt"

type person struct {
	name string
}

func main() {
	p := person{
		name: "nhan",
	}

	p.setName("aa")
	fmt.Println(p)
}

func (p *person) setName(n string) {
	p.name = n
}
