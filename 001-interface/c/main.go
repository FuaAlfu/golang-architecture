package main

import(
	"fmt"
)

type(
	Person struct{
		name string
	}

	accessor interface{
		save(n int,p Person)
		retrieve(n int)Person
	}

	personService struct{
		a accessor
	}

	mongo   map[int]Person
	postger map[int]Person
)

//recieve a person from our database..
func (ps personService) get(n int)(Person, error){
	p:= ps.a.retrieve(n)
	if p.name == ""{
		return Person{}, fmt.Errorf("no person with n of %d", n)
	}
	return p, nil
}

func (m mongo) save(n int, p Person){
	m[n] = p
}

//retrieve
func (m mongo) retrieve(n int) Person{
	return m[n]
}

func (pg postger) save(n int, p Person){
	pg[n] = p
}

//retrieve
func (pg postger) retrieve(n int) Person{
	return pg[n]
}

func put(a accessor, n int, p Person){
	a.save(n,p)
}

func get(a accessor, n int) Person{
	return a.retrieve(n)
}

func main() {
	//create database
	dbm  := mongo{}
	dbpg := postger{}

	//costructing
	ps := personService{
		a: dbm,
	}
	p1 := Person{
		name: "john",
	}

	p2 := Person{
		name: "ada",
	}


	//ps.get()
	put(dbm, 1, p1)
	put(dbm, 2, p2)
	fmt.Println(get(dbm,1))
	fmt.Println(get(dbm,2))

	fmt.Println(ps.get(1))
	fmt.Println(ps.get(3)) //error ..

	//-----------(other data storage)
	put(dbpg, 1, p1)
	put(dbpg, 2, p2)
	fmt.Println(get(dbpg,1))
	fmt.Println(get(dbpg,2))
}