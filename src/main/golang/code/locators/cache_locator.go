package locators

import (
	"github.com/starter-go/env"
)

// CacheLocator ...
type CacheLocator struct {

	//starter:component

	_as func(env.LocatorRegistry) //starter:as(".")

}

func (inst *CacheLocator) _impl() (env.LocatorRegistry, env.Locator) {
	return inst, inst
}

// Registrations ...
func (inst *CacheLocator) Registrations() []*env.LocatorRegistration {
	lr := &env.LocatorRegistration{
		Locator:  inst,
		Enabled:  true,
		Priority: PriorityCache,
	}
	return []*env.LocatorRegistration{lr}
}

// Locate ...
func (inst *CacheLocator) Locate(want *env.Query, chain env.LocatorChain) error {
	return chain.Locate(want)
}
