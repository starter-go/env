package locators

import (
	"github.com/starter-go/env"
)

// ExampleLocator ...
type ExampleLocator struct {

	//starter:component

	_as func(env.LocatorRegistry) //starter:as(".")

}

func (inst *ExampleLocator) _impl() (env.LocatorRegistry, env.Locator) {
	return inst, inst
}

// Registrations ...
func (inst *ExampleLocator) Registrations() []*env.LocatorRegistration {
	lr := &env.LocatorRegistration{
		Locator:  inst,
		Enabled:  true,
		Priority: PriorityExample,
	}
	return []*env.LocatorRegistration{lr}
}

// Locate ...
func (inst *ExampleLocator) Locate(want *env.Query, chain env.LocatorChain) error {
	return chain.Locate(want)
}
