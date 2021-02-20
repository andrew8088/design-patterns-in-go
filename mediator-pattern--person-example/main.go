package main

func main() {
	//m := NewArrayMediator()
	m := NewListMediator()

	ann := NewPerson("Ann", m)
	bob := NewPerson("Bob", m)

	NewPerson("Chris", m)
	NewPerson("Dani", m)

	ann.send("Hello world!")
	bob.send("Goodbye world!")
}
