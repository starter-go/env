package locators

import (
	"os/user"

	"github.com/starter-go/afs"
	"github.com/starter-go/env"

	"strings"
)

// DefaultLocator ... 把所有的请求统一定位到 "~/.appname/"
type DefaultLocator struct {

	//starter:component

	_as func(env.LocatorRegistry) //starter:as(".")

	FS afs.FS //starter:inject("#")

}

func (inst *DefaultLocator) _impl() (env.LocatorRegistry, env.Locator) {
	return inst, inst
}

// Registrations ...
func (inst *DefaultLocator) Registrations() []*env.LocatorRegistration {
	lr := &env.LocatorRegistration{
		Locator:  inst,
		Enabled:  true,
		Priority: PriorityDefault,
	}
	return []*env.LocatorRegistration{lr}
}

// Locate ...
func (inst *DefaultLocator) Locate(q *env.Query, chain env.LocatorChain) error {
	appname := strings.ToLower(q.Want.App)
	home, err := inst.getUserHome()
	if err != nil {
		return err
	}
	target := home.GetChild("." + appname + "/default")
	q.Have.Path = target
	return nil
}

func (inst *DefaultLocator) getUserHome() (afs.Path, error) {
	u, err := user.Current()
	if err != nil {
		return nil, err
	}
	home := inst.FS.NewPath(u.HomeDir)
	return home, nil
}
