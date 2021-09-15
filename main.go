package main

import "fmt"

type person struct{
	name string
	age int
}


func newPerson(name string)*person {
	p:= person{name:name}	
	p.age = 20

	return &p
}


func main() {
	fmt.Println(person{"bob", 20}) 
	fmt.Println(person{name: "sakura endo", age: 20})
	fmt.Println(person{name:"miku"})
	fmt.Println(&person{name: "anin", age: 20})
	fmt.Println(newPerson("andini"))

	s := person{name: "sakura miyawaki", age: 22}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)
	sp.age = 23
	fmt.Println(sp.age)
}

	



