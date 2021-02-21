package main

type Service interface {
	// initialized bool
	process(event string, metadata string)
}
