package resolvers

import (
	"github.com/starter-go/env"
)

// CacheResolver ...
type CacheResolver struct {

	//starter:component

	_as func(env.ResolverRegistry) //starter:as(".")

}

func (inst *CacheResolver) _impl() (env.ResolverRegistry, env.Resolver) {
	return inst, inst
}

// Registrations ...
func (inst *CacheResolver) Registrations() []*env.ResolverRegistration {
	lr := &env.ResolverRegistration{
		Resolver: inst,
		Enabled:  true,
		Priority: PriorityCache,
	}
	return []*env.ResolverRegistration{lr}
}

// Resolve ...
func (inst *CacheResolver) Resolve(want *env.Query, chain env.ResolverChain) error {

	// todo ...
	return chain.Resolve(want)
}
