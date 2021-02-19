package main

type ArrayMediator struct {
	people *[10]Person
}

func (m ArrayMediator) Send(message string, originator Person) {
	for i := 0; i < len(m.people); i++ {
		person := m.people[i]
		if person != originator && person.name != "" {
			person.receive(message)
		}
	}
}

func (m ArrayMediator) AddPerson(person Person) {
	for i := 0; i < len(m.people); i++ {
		if m.people[i].name == "" {
			m.people[i] = person;
			return
		}
	}
}

func NewArrayMediator() Mediator {
	a := [10]Person{}
	return ArrayMediator{&a}
}
