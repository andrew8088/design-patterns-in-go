package main

type ArrayEventMediator struct {
	services *[10]Service
}

func (m ArrayEventMediator) fire(event string, metadata string) {
	for i := 0; i < len(m.services); i++ {
		if m.services[i] != nil {
			m.services[i].process(event, metadata)
		}
	}
}

func (m ArrayEventMediator) registerService(service Service) {
	for i := 0; i < len(m.services); i++ {
		if m.services[i] == nil {
			m.services[i] = service;
			return
		}
	}
}

func NewArrayEventMediator() ArrayEventMediator {
	a := [10]Service{}
	return ArrayEventMediator{&a}
}
