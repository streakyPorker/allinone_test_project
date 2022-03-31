package oop

import (
	"fmt"
	"reflect"
)

func init() {
	fmt.Println("init oop")
}

type Person struct {
	name string
	age  int
}

type Person2 struct {
	age2 int
}

type Worker interface {
	work(id int) (desc string)
	rest()
}

type ITWorker interface {
	Worker
	workWithComputer()
}

type SW interface {
	Ref(int) int
	Ptr(int) int
}

type SkilledWorker struct {
	Person
	Person2
}

func (p SkilledWorker) Ref(i int) int {
	return i + 1
}

func (p *SkilledWorker) Ptr(i int) int {
	return i + 2
}

func (p Person) work(id int) (desc string) {

	fmt.Println("id is", id, "info is ", p)
	return fmt.Sprintf(p.name+"`s id is %d", id)
}

func (p Person) rest() {
	fmt.Println(p.name, "is resting")
}

func (p SkilledWorker) rest() {
	fmt.Println(p.name, "is resting skillfully")
}

func (p Person) workWithComputer() {
	fmt.Println(p.name, "is working with computer")
}

func TestFunc() {
	sw := SkilledWorker{}
	sw.name = "init"

	person := Person{}
	fmt.Println(sw, person)
	var itWorker ITWorker = &person
	person.name = "personName"
	worker := itWorker.(*Person)
	fmt.Println(worker, reflect.TypeOf(worker))
	worker.name = "givenName"
	worker.age = 22
	fmt.Println(person.name, worker.name)
	sw.Person = person
	sw.Person.work(12)

	switch v := itWorker.(type) {
	case nil:
		fmt.Println("nil", v)
	default:
		fmt.Println(reflect.TypeOf(sw), v)
	}

}

func InterfaceTestFunc() {
}

func isWorker(worker Worker) {
	fmt.Println(worker.work(12))
}
