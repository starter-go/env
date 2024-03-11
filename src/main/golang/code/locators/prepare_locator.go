package locators

import (
	"sync"

	"github.com/starter-go/env"
)

// PrepareLocator ...
type PrepareLocator struct {

	//starter:component

	_as func(env.LocatorRegistry) //starter:as(".")

	SimpleAppName string //starter:inject("${application.simple-name}")

	mutex sync.Mutex
}

func (inst *PrepareLocator) _impl() (env.LocatorRegistry, env.Locator) {
	return inst, inst
}

// Registrations ...
func (inst *PrepareLocator) Registrations() []*env.LocatorRegistration {
	lr := &env.LocatorRegistration{
		Locator:  inst,
		Enabled:  true,
		Priority: PriorityPrepare,
	}
	return []*env.LocatorRegistration{lr}
}

// Locate ...
func (inst *PrepareLocator) Locate(q *env.Query, chain env.LocatorChain) error {

	inst.prepare(q.Want)
	inst.mutex.Lock()
	defer func() {
		inst.mutex.Unlock()
	}()

	return chain.Locate(q)
}

func (inst *PrepareLocator) prepare(want *env.Want) {

	if want.App == "" {
		want.App = inst.SimpleAppName
	}

	if want.Env == "" {
		want.Env = env.ComputeEnvName(want)
	}

	if want.Property == "" {
		want.Property = env.ComputePropertyName(want)
	}
}
