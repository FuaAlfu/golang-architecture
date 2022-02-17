package Postg

import "github.com/FuaAlfu/golang-architecture.git"

type Postger map[int]architecture.Person

func (pg Postger) Save(n int, p architecture.Person){
	pg[n] = p
}

//retrieve
func (pg Postger) Retrieve(n int) architecture.Person{
	return pg[n]
}