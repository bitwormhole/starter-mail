package main

import (
	"github.com/bitwormhole/starter"
	startermail "github.com/bitwormhole/starter-mail"
)

func main() {

	i := starter.InitApp()
	i.Use(startermail.Module())
	i.Run()

}
