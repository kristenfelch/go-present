package create

import (
	"github.com/kristenfelch/go-present/demo"
	"github.com/kristenfelch/go-present/monitor"
)

func create() {
	demo.Decorate(CreateArrayRunner{}, monitor.Monitor("CreateArrayRunner")).Run()
	demo.Decorate(UpdateArrayRunner{}, monitor.Monitor("UpdateArrayRunner")).Run()
	demo.Decorate(CreateStringRunner{}, monitor.Monitor("CreateStringRunner")).Run()
	demo.Decorate(UpdateStringRunner{}, monitor.Monitor("UpdateStringRunner")).Run()
}

type CreateArrayRunner struct{}

func (CreateArrayRunner) Run() {
	for i := 0; i < 1000000; i++ {
		languages := [4]string{"Java", "Scala"}
		languages[2] = "Node"
		languages[3] = "Golang"
	}
}

type UpdateArrayRunner struct{}

func (UpdateArrayRunner) Run() {
	for i := 0; i < 1000000; i++ {
		languages := []string{"Java", "Scala"}
		_ = append(languages, "Node", "Golang")
	}
}

type CreateStringRunner struct{}

func (CreateStringRunner) Run() {
	for i := 0; i < 1000000; i++ {
		name := "first"
		name = name + " last"
	}
}

type UpdateStringRunner struct{}

func (UpdateStringRunner) Run() {
	for i := 0; i < 1000000; i++ {
		first := "first"
		last := "last"
		_ = first + last
	}
}
