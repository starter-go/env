package locators

import (
	"github.com/starter-go/env"
)

// PropertyBasedLocator ...
type PropertyBasedLocator struct {

	//starter:component

	_as func(env.LocatorRegistry) //starter:as(".")

}

func (inst *PropertyBasedLocator) _impl() (env.LocatorRegistry, env.Locator) {
	return inst, inst
}

// Registrations ...
func (inst *PropertyBasedLocator) Registrations() []*env.LocatorRegistration {
	lr := &env.LocatorRegistration{
		Locator:  inst,
		Enabled:  true,
		Priority: PriorityPropertyBased,
	}
	return []*env.LocatorRegistration{lr}
}

// Locate ...
func (inst *PropertyBasedLocator) Locate(want *env.Query, chain env.LocatorChain) error {
	return chain.Locate(want)
}
