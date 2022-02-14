package main

import (
	"github.com/bitwormhole/starter"
	startermail "github.com/bitwormhole/starter-mail"
)

func main() {
	err := run()
	if err != nil {
		panic(err)
	}
}

func run() error {
	i := starter.InitApp()
	i.Use(startermail.ModuleForDemo())
	rt, err := i.RunEx()
	if err != nil {
		return err
	}
	return rt.Exit()
}
