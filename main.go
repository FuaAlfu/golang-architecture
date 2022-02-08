package main

import "fmt"

type(
	Person struct{
		//concrete
		name string
	}
	Human interface{
		//abstract
		speak()
	}
)

var x Human

func (p Person) speak(){
	fmt.Println("from a person - this is my name: ",p.name)
}

func main() {
	p1 := Person{
		name: "john",
	}

	//speak(p1.name)
	//---
	x = p1
	fmt.Printf("%T\n",x)
}