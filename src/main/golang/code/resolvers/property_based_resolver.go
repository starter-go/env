package resolvers

import (
	"github.com/starter-go/application"
	"github.com/starter-go/application/properties"
	"github.com/starter-go/env"
)

// PropertyBasedResolver ...
type PropertyBasedResolver struct {

	//starter:component

	_as func(env.ResolverRegistry) //starter:as(".")

	AppContext application.Context //starter:inject("context")

	values properties.Table
}

func (inst *PropertyBasedResolver) _impl() (env.ResolverRegistry, env.Resolver) {
	return inst, inst
}

// Registrations ...
func (inst *PropertyBasedResolver) Registrations() []*env.ResolverRegistration {
	lr := &env.ResolverRegistration{
		Resolver: inst,
		Enabled:  true,
		Priority: PriorityPropertyBased,
	}
	return []*env.ResolverRegistration{lr}
}

// Resolve ...
func (inst *PropertyBasedResolver) Resolve(query *env.Query, chain env.ResolverChain) error {
	name := query.Want.Property
	if name != "" {
		table := inst.getTable()
		value := table.GetProperty(name)
		if value != "" {
			query.Have.Value = value
			return nil
		}
	}
	return chain.Resolve(query)
}

func (inst *PropertyBasedResolver) getTable() properties.Table {
	table := inst.values
	if table == nil {
		table = inst.AppContext.GetProperties()
		inst.values = table
	}
	return table
}
