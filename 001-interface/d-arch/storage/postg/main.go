package Postg

import "github.com/FuaAlfu/golang-architecture.git"

type Postger map[int]architecture.Person

func (pg Postger) save(n int, p architecture.Person){
	pg[n] = p
}

//retrieve
func (pg Postger) retrieve(n int) architecture.Person{
	return pg[n]
}