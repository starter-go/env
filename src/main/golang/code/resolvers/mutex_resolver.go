package resolvers

import (
	"sync"

	"github.com/starter-go/env"
)

// MutexResolver ...
type MutexResolver struct {

	//starter:component

	_as func(env.ResolverRegistry) //starter:as(".")

	mutex sync.Mutex
}

func (inst *MutexResolver) _impl() (env.ResolverRegistry, env.Resolver) {
	return inst, inst
}

// Registrations ...
func (inst *MutexResolver) Registrations() []*env.ResolverRegistration {
	lr := &env.ResolverRegistration{
		Resolver: inst,
		Enabled:  true,
		Priority: PriorityMutex,
	}
	return []*env.ResolverRegistration{lr}
}

// Resolve ...
func (inst *MutexResolver) Resolve(q *env.Query, chain env.ResolverChain) error {
	inst.mutex.Lock()
	defer func() {
		inst.mutex.Unlock()
	}()
	return chain.Resolve(q)
}
