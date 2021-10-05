package main

import "fmt"

type person struct {
	Name string
	Age int
	Gender string
}

type OptionFunc func(*person)

func withName(n string) OptionFunc {
	return func(p *person) {
		p.Name = n
	}
}

func withAge(a int) OptionFunc {
	return func(p *person) {
		p.Age = a
	}
}

func withGender(g string) OptionFunc {
	return func(p *person) {
		p.Gender = g
	}
}

var defaultPerson = &person{
	"why",
	18,
	"man",
}

func initPerson(opts ...OptionFunc) (p *person) {
	p = defaultPerson
	for _, o := range opts {
		o(p)
	}
	return
}

func main() {
	p1 := initPerson()
	fmt.Println(p1)
	p2 := initPerson(
		withName("superyong"))
	fmt.Println(p2)

}
