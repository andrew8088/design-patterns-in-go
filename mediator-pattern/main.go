package main

func main() {
	mediator := NewArrayEventMediator()
	registrationService := NewRegistrationService(mediator)

	NewEmailService(mediator)

	registrationService.RegisterMember("anna")
	registrationService.RegisterMember("bob")
}
