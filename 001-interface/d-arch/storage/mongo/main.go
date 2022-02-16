package mongo

import "github.com/FuaAlfu/golang-architecture.git"

type Mongo   map[int]architecture.Person

func (m Mongo) save(n int, p .architecturePerson){
	m[n] = p
}

//retrieve
func (m Mongo) retrieve(n int) architecture.Person{
	return m[n]
}