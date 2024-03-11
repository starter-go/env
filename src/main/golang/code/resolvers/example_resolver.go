package resolvers

import (
	"github.com/starter-go/env"
)

// ExampleResolver ...
type ExampleResolver struct {

	//starter:component

	_as func(env.ResolverRegistry) //starter:as(".")

}

func (inst *ExampleResolver) _impl() (env.ResolverRegistry, env.Resolver) {
	return inst, inst
}

// Registrations ...
func (inst *ExampleResolver) Registrations() []*env.ResolverRegistration {
	lr := &env.ResolverRegistration{
		Resolver: inst,
		Enabled:  true,
		Priority: PriorityExample,
	}
	return []*env.ResolverRegistration{lr}
}

// Resolve ...
func (inst *ExampleResolver) Resolve(want *env.Query, chain env.ResolverChain) error {
	return chain.Resolve(want)
}
