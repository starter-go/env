package code

import (
	"fmt"

	"github.com/starter-go/afs"
	"github.com/starter-go/application/environment"
	"github.com/starter-go/application/properties"
	"github.com/starter-go/env"
)

// EnvServiceImpl ...
type EnvServiceImpl struct {

	//starter:component

	_as func(env.Service) //starter:as("#")

	Resolvers []env.ResolverRegistry //starter:inject(".")

	chain env.ResolverChain
}

func (inst *EnvServiceImpl) _impl() env.Service {
	return inst
}

// ForApp ...
func (inst *EnvServiceImpl) ForApp(appName string) env.Context {
	ctx := new(envAppCtx)
	ctx.appName = appName
	ctx.chain = inst.getResolverChain()
	return ctx
}

func (inst *EnvServiceImpl) getResolverChain() env.ResolverChain {
	chain := inst.chain
	if chain == nil {
		chain = inst.loadResolverChain()
		inst.chain = chain
	}
	return chain
}

func (inst *EnvServiceImpl) loadResolverChain() env.ResolverChain {
	src := inst.Resolvers
	builder := new(resolverChainBuilder)
	for _, r1 := range src {
		r2s := r1.Registrations()
		builder.add(r2s...)
	}
	return builder.create()
}

////////////////////////////////////////////////////////////////////////////////

type envAppCtx struct {
	chain   env.ResolverChain
	appName string
}

func (inst *envAppCtx) _impl() env.Context {
	return inst
}

func (inst *envAppCtx) AppName() string {
	return inst.appName
}

// Query ...
func (inst *envAppCtx) Query(want *env.Want) (*env.Have, error) {
	q := &env.Query{}
	q.Want = want
	err := inst.doQuery(q)
	if err != nil {
		return nil, err
	}
	return q.Have, nil
}

func (inst *envAppCtx) doQuery(q *env.Query) error {
	if q == nil {
		return fmt.Errorf("no param:query")
	}
	if q.Want == nil {
		q.Want = &env.Want{}
	}
	if q.Have == nil {
		q.Have = &env.Have{}
	}
	q.Want.App = inst.appName
	q.Have.Want = q.Want
	chain := inst.chain
	err := chain.Resolve(q)
	if err != nil {
		return err
	}
	if q.Have.Path == nil {
		panic("env.Have.Path is nil")
	}
	return nil
}

// GetPath ...
func (inst *envAppCtx) GetPath(q *env.Want) afs.Path {
	have, err := inst.Query(q)
	if err != nil {
		panic(err)
	}
	return have.Path
}

// GetEnvironment ...
func (inst *envAppCtx) GetEnvironment() environment.Table {
	dst := environment.NewTable(nil)
	q := &env.Query{}
	q.Names = make(map[string]*env.Want)
	err := inst.doQuery(q)
	if err != nil {
		return dst
	}
	src := q.Names
	for _, item := range src {
		have, err := inst.Query(item)
		if err == nil && have != nil {
			dst.SetEnv(item.Env, have.Value)
		}
	}
	return dst
}

// GetProperties ...
func (inst *envAppCtx) GetProperties() properties.Table {
	dst := properties.NewTable(nil)
	q := &env.Query{}
	q.Names = make(map[string]*env.Want)
	err := inst.doQuery(q)
	if err != nil {
		return dst
	}
	src := q.Names
	for _, item := range src {
		have, err := inst.Query(item)
		if err == nil && have != nil {
			dst.SetProperty(item.Property, have.Value)
		}
	}
	return dst
}
