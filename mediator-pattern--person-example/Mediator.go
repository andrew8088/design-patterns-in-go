package main

type Mediator interface {
	Send(message string, sender Person)
	AddPerson(person Person)
}
