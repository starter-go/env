package code

import (
	"github.com/starter-go/afs"
	"github.com/starter-go/application/environment"
	"github.com/starter-go/application/properties"
	"github.com/starter-go/env"
)

// EnvServiceImpl ...
type EnvServiceImpl struct {

	//starter:component

	_as func(env.Service) //starter:as("#")

	Locators []env.LocatorRegistry //starter:inject(".")

	chain env.LocatorChain
}

func (inst *EnvServiceImpl) _impl() env.Service {
	return inst
}

// ForApp ...
func (inst *EnvServiceImpl) ForApp(appName string) env.Context {
	ctx := new(envAppCtx)
	ctx.appName = appName
	ctx.chain = inst.getLocatorChain()
	return ctx
}

func (inst *EnvServiceImpl) getLocatorChain() env.LocatorChain {
	chain := inst.chain
	if chain == nil {
		chain = inst.loadLocatorChain()
		inst.chain = chain
	}
	return chain
}

func (inst *EnvServiceImpl) loadLocatorChain() env.LocatorChain {
	src := inst.Locators
	builder := new(locatorChainBuilder)
	for _, r1 := range src {
		r2s := r1.Registrations()
		builder.add(r2s...)
	}
	return builder.create()
}

////////////////////////////////////////////////////////////////////////////////

type envAppCtx struct {
	chain   env.LocatorChain
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
	want.App = inst.appName
	have := &env.Have{Want: want}
	q := &env.Query{Want: want, Have: have}
	chain := inst.chain
	err := chain.Locate(q)
	if err != nil {
		return nil, err
	}
	if have.Path == nil {
		panic("env.Have.Path is nil")
	}
	return have, nil
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
	src := env.ListEnvNames(inst.appName)
	for _, name := range src {
		want := &env.Want{Env: name}
		path := inst.GetPath(want)
		if path != nil {
			value := path.GetPath()
			dst.SetEnv(name, value)
		}
	}
	return dst
}

// GetProperties ...
func (inst *envAppCtx) GetProperties() properties.Table {
	dst := properties.NewTable(nil)
	src := env.ListPropertyNames(inst.appName)
	for _, name := range src {
		want := &env.Want{Property: name}
		path := inst.GetPath(want)
		if path != nil {
			value := path.GetPath()
			dst.SetProperty(name, value)
		}
	}
	return dst
}
