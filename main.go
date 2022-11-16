package main

import "fmt"

type Message string

type Greeter struct {
	Message Message
}

type Event struct {
	Greeter Greeter // <- adding a Greeter field
}

func NewMessage() Message {
	return Message("Hi there!")
}

func NewGreeter(m Message) Greeter {
	return Greeter{Message: m}
}

func (g Greeter) Greet() Message {
	return g.Message
}

func NewEvent(g Greeter) Event {
	return Event{Greeter: g}
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

// func main() {
//     message := NewMessage()
//     greeter := NewGreeter(message)
//     event := NewEvent(greeter) //DI를 구현하는데 초기화 과정이 너무많이듬

//     event.Start()
// }

func main() {
	e := InitializeEvent()

	e.Start()
}
