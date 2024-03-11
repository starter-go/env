package main

import (
	"os"

	"github.com/starter-go/env/modules/env"
	"github.com/starter-go/units"
)

func main() {
	m := env.ModuleForTest()

	// i := starter.Init(os.Args)
	// i.MainModule(m)
	// i.WithPanic(true).Run()

	ur := units.NewRunner()
	ur.Dependencies(m)
	ur.EnablePanic(true).Run(os.Args)
}
