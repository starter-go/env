package resolvers

import (
	"strings"

	"github.com/starter-go/env"
)

// NamesResolver ...
type NamesResolver struct {

	//starter:component

	_as func(env.ResolverRegistry) //starter:as(".")

}

func (inst *NamesResolver) _impl() (env.ResolverRegistry, env.Resolver) {
	return inst, inst
}

// Registrations ...
func (inst *NamesResolver) Registrations() []*env.ResolverRegistration {
	lr := &env.ResolverRegistration{
		Resolver: inst,
		Enabled:  true,
		Priority: PriorityNames,
	}
	return []*env.ResolverRegistration{lr}
}

// Resolve ...
func (inst *NamesResolver) Resolve(q *env.Query, chain env.ResolverChain) error {

	names := q.Names
	if names != nil {
		src := inst.listNames(q)
		for _, item := range src {
			names[item.Env] = item
		}
	}

	return chain.Resolve(q)
}

func (inst *NamesResolver) listNames(q *env.Query) []*env.Want {
	appName := q.Want.App
	dst := make([]*env.Want, 0)
	src := env.ListEnvNames(appName)
	for _, name := range src {
		pName := strings.ToLower(name)
		pName = strings.ReplaceAll(pName, "_", ".")
		item := &env.Want{
			Env:      name,
			Property: pName,
			App:      appName,
		}
		dst = append(dst, item)
	}
	return dst
}
