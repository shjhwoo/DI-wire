package main

func main() {
	e := InitializeEvent()

	e.Start()
}

// wire.go

func InitializeEvent() Event {
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{}
}