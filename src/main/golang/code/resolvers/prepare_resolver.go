package resolvers

import (
	"github.com/starter-go/afs"
	"github.com/starter-go/env"
)

// PrepareResolver ...
type PrepareResolver struct {

	//starter:component

	_as func(env.ResolverRegistry) //starter:as(".")

	SimpleAppName string //starter:inject("${application.simple-name}")
	FS            afs.FS //starter:inject("#")

}

func (inst *PrepareResolver) _impl() (env.ResolverRegistry, env.Resolver) {
	return inst, inst
}

// Registrations ...
func (inst *PrepareResolver) Registrations() []*env.ResolverRegistration {
	lr := &env.ResolverRegistration{
		Resolver: inst,
		Enabled:  true,
		Priority: PriorityPrepare,
	}
	return []*env.ResolverRegistration{lr}
}

// Resolve ...
func (inst *PrepareResolver) Resolve(q *env.Query, chain env.ResolverChain) error {
	inst.prepare(q.Want)
	err := chain.Resolve(q)
	if err == nil {
		inst.handlePost(q)
	}
	return err
}

func (inst *PrepareResolver) handlePost(q *env.Query) {

	path := q.Have.Path
	value := q.Have.Value

	if path == nil && value == "" {
		// nop
	} else if path != nil {
		q.Have.Value = path.GetPath()
	} else if value != "" {
		// try as path
		path = inst.FS.NewPath(value)
		if path.Exists() {
			q.Have.Path = path
		}
	}
}

func (inst *PrepareResolver) prepare(want *env.Want) {

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
