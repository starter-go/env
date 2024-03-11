package resolvers

import (
	"os/user"

	"github.com/starter-go/afs"
	"github.com/starter-go/env"

	"strings"
)

// DefaultResolver ... 把所有的请求统一定位到 "~/.appname/default"
type DefaultResolver struct {

	//starter:component

	_as func(env.ResolverRegistry) //starter:as(".")

	FS afs.FS //starter:inject("#")

}

func (inst *DefaultResolver) _impl() (env.ResolverRegistry, env.Resolver) {
	return inst, inst
}

// Registrations ...
func (inst *DefaultResolver) Registrations() []*env.ResolverRegistration {
	lr := &env.ResolverRegistration{
		Resolver: inst,
		Enabled:  true,
		Priority: PriorityDefault,
	}
	return []*env.ResolverRegistration{lr}
}

// Resolve ...
func (inst *DefaultResolver) Resolve(q *env.Query, chain env.ResolverChain) error {
	appname := strings.ToLower(q.Want.App)
	home, err := inst.getUserHome()
	if err != nil {
		return err
	}
	target := home.GetChild("." + appname + "/default")
	q.Have.Path = target
	return nil
}

func (inst *DefaultResolver) getUserHome() (afs.Path, error) {
	u, err := user.Current()
	if err != nil {
		return nil, err
	}
	home := inst.FS.NewPath(u.HomeDir)
	return home, nil
}
