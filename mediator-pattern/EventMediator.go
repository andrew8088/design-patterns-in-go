package main

type EventMediator interface {
	fire(event string, metadata string)
	registerService(service Service)
}
