//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

// wire.go

func InitializeEvent() Event {
	//wire.Build는 wire 명령어를 통해 자동으로 초기 생성자를 만들도록 한다.
	//build안에 전달하는 인수를 provider라고 한다.
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{}
}
