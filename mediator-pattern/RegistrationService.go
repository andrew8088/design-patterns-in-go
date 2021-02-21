package main

import "fmt"

type RegistrationService struct {
	eventMediator EventMediator
}

func (s RegistrationService) process(event string, metadata string) {
	if event == "activation-email-sent" {
		fmt.Println("RegistrationService - activation email sent to", metadata);
	}

	if event == "activation-email-failed" {
		fmt.Println("RegistrationService - failed to send activation email to", metadata);
	}
}

func (s RegistrationService) RegisterMember(name string) {
	fmt.Println("RegistrationService - registered member", name)
	s.eventMediator.fire("member-registered", name)
}

func NewRegistrationService(eventMediator EventMediator) RegistrationService {
	r := RegistrationService{ eventMediator }
	eventMediator.registerService(r)
	return r
}
