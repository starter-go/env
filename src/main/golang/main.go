package main

import (
	"os"

	"github.com/starter-go/env/modules/env"
	"github.com/starter-go/starter"
)

func main() {
	m := env.Module()
	i := starter.Init(os.Args)
	i.MainModule(m)
	i.WithPanic(true).Run()
}
