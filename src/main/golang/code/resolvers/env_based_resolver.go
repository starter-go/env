package resolvers

import (
	"github.com/starter-go/application"
	"github.com/starter-go/application/environment"
	"github.com/starter-go/env"
)

// EnvBasedResolver ...
type EnvBasedResolver struct {

	//starter:component

	_as func(env.ResolverRegistry) //starter:as(".")

	AppContext application.Context //starter:inject("context")

	values environment.Table
}

func (inst *EnvBasedResolver) _impl() (env.ResolverRegistry, env.Resolver) {
	return inst, inst
}

// Registrations ...
func (inst *EnvBasedResolver) Registrations() []*env.ResolverRegistration {
	lr := &env.ResolverRegistration{
		Resolver: inst,
		Enabled:  true,
		Priority: PriorityEnvBased,
	}
	return []*env.ResolverRegistration{lr}
}

// Resolve ...
func (inst *EnvBasedResolver) Resolve(query *env.Query, chain env.ResolverChain) error {
	name := query.Want.Env
	if name != "" {
		table := inst.getTable()
		value := table.GetEnv(name)
		if value != "" {
			query.Have.Value = value
			return nil
		}
	}
	return chain.Resolve(query)
}

func (inst *EnvBasedResolver) getTable() environment.Table {
	table := inst.values
	if table == nil {
		table = inst.AppContext.GetEnvironment()
		inst.values = table
	}
	return table
}
