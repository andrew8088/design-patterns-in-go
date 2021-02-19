package main

import (
	"container/list"
	"time"
)

type SlowMediator struct {
	people *list.List
}

func (m SlowMediator) send(message string, originator Person) {
	for element := m.people.Front(); element != nil; element = element.Next() {
		time.Sleep(time.Second)
		person := element.Value
		if person != originator {
			person.(Person).receive(message)
		}
	}
}

func (m SlowMediator) addPerson(person Person) {
	m.people.PushFront(person)
}
