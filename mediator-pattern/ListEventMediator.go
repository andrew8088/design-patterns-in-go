package main

import "container/list"

type ListEventMediator struct {
	services *list.List
}

func (m ListEventMediator) fire(event string, metadata string) {
	for element := m.services.Front(); element != nil; element = element.Next() {
		element.Value.(Service).process(event, metadata)
	}
}

func (m ListEventMediator) registerService(service Service) {
	m.services.PushFront(service)
}

func NewListEventMediator() EventMediator {
	l := list.New()
	return ListEventMediator{l}
}
