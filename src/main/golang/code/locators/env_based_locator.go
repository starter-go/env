package locators

import (
	"github.com/starter-go/env"
)

// EnvBasedLocator ...
type EnvBasedLocator struct {

	//starter:component

	_as func(env.LocatorRegistry) //starter:as(".")

}

func (inst *EnvBasedLocator) _impl() (env.LocatorRegistry, env.Locator) {
	return inst, inst
}

// Registrations ...
func (inst *EnvBasedLocator) Registrations() []*env.LocatorRegistration {
	lr := &env.LocatorRegistration{
		Locator:  inst,
		Enabled:  true,
		Priority: PriorityEnvBased,
	}
	return []*env.LocatorRegistration{lr}
}

// Locate ...
func (inst *EnvBasedLocator) Locate(want *env.Query, chain env.LocatorChain) error {
	return chain.Locate(want)
}
