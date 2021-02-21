package main

import "fmt"

type EmailService struct {
	eventMediator EventMediator
}

func (s EmailService) process(event string, metadata string) {
	if event == "member-registered" {
		if metadata == "anna" {
			fmt.Println("EmailService - sent activation email to", metadata);
			s.eventMediator.fire("activation-email-sent", metadata)
		} else {
			fmt.Println("EmailService - failed to send to", metadata);
			s.eventMediator.fire("activation-email-failed", metadata)
		}
	}
}

func NewEmailService(eventMediator EventMediator) EmailService {
	r := EmailService{ eventMediator }
	eventMediator.registerService(r)
	return r
}
