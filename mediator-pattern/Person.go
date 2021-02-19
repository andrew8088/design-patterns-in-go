package main

import "fmt"

type Person struct {
	name     string
	mediator Mediator
}

func (p Person) receive(message string) {
	fmt.Println(p.name + " received message: " + message)
}

func (p Person) send(message string) {
	p.mediator.Send(message, p)
}

func NewPerson(name string, mediator Mediator) Person {
	p := Person{name, mediator}
	mediator.AddPerson(p)
	return p
}

