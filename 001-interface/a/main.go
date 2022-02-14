package main

import "fmt"

type(
	Person struct{
		//concrete
		name string
	}
	SecretAgent struct{
	    Person
		ltk bool
	}
	Human interface{
		//abstract
		speak()
	}
)

var x,y Human

func (p Person) speak(){
	fmt.Println("from a person - this is my name: ",p.name)
}

func (sa SecretAgent) speak(){
	fmt.Println("I'm a secret agent and - this is my name: ",sa.name)
}

func foo(h Human){
	h.speak()
}

func main() {
	p1 := Person{
		name: "john",
	}

	p2 := SecretAgent{
		Person: Person{
			name: "doe",
		},
		ltk: true,
	}

	//speak(p1.name)
	//---
	x = p1
	x.speak()

	y = p2
	y.speak()
	fmt.Printf("%T\n",x)
	fmt.Printf("%T\n",y)
	println("---")
	foo(x)
	foo(y)
	foo(p1)
	foo(p2)
}