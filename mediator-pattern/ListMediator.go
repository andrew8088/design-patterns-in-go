 package main

import "container/list"

type ListMediator struct {
	people *list.List
}

func (m ListMediator) Send(message string, originator Person) {
	for element := m.people.Front(); element != nil; element = element.Next() {
		person := element.Value
		if person != originator {
			person.(Person).receive(message)
		}
	}
}

func (m ListMediator) AddPerson(person Person) {
	m.people.PushFront(person)
}

func NewListMediator() Mediator {
	l := list.New()
	return ListMediator{l}
}
