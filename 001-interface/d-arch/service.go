package architecture

import(
	"fmt"
)

type(
	Person struct{
		Name string
	}

	Accessor interface{
		Save(n int,p Person)
		Retrieve(n int)Person
	}

	PersonService struct{
		a Accessor
	}
)

//recieve a person from our database..
func (ps PersonService) get(n int)(Person, error){
	p:= ps.a.Retrieve(n)
	if p.Name == ""{
		return Person{}, fmt.Errorf("no person with n of %d", n)
	}
	return p, nil
}

func put(a Accessor, n int, p Person){
	a.Save(n,p)
}

func get(a Accessor, n int) Person{
	return a.Retrieve(n)
}

func main() {
	//create database
	dbm  := mongo{}
	dbpg := postger{}

	//costructing
	ps := PersonService{
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