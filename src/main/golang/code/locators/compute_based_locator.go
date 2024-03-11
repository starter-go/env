package locators

import (
	"fmt"
	"os"
	"os/user"
	"runtime"
	"strings"

	"github.com/starter-go/afs"
	"github.com/starter-go/env"
)

const (
	envRegularPrefix = "REG_"

	envAppHome    = envRegularPrefix + "HOME"
	envAppData    = envRegularPrefix + "DATA"
	envAppCode    = envRegularPrefix + "CODE"
	envAppConfig  = envRegularPrefix + "CONFIG"
	envAppSetting = envRegularPrefix + "SETTING"

	envUserAppHome = envRegularPrefix + "USER_APP_HOME"
	envUserHome    = envRegularPrefix + "USER_HOME"

	envShareSystemHome = envRegularPrefix + "SHARE_SYSTEM_HOME"
)

// ComputeBasedLocator ...
type ComputeBasedLocator struct {

	//starter:component

	_as func(env.LocatorRegistry) //starter:as(".")

	FS afs.FS //starter:inject("#")

	handlers map[string]func(q *env.Query) error
}

func (inst *ComputeBasedLocator) _impl() (env.LocatorRegistry, env.Locator) {
	return inst, inst
}

// Registrations ...
func (inst *ComputeBasedLocator) Registrations() []*env.LocatorRegistration {
	lr := &env.LocatorRegistration{
		Locator:  inst,
		Enabled:  true,
		Priority: PriorityComputeBased,
	}
	return []*env.LocatorRegistration{lr}
}

// Locate ...
func (inst *ComputeBasedLocator) Locate(query *env.Query, chain env.LocatorChain) error {

	err := inst.tryLocateAsEnv(query)
	if err == nil {
		return nil
	}

	err = inst.tryLocateAsProperty(query)
	if err == nil {
		return nil
	}

	// return err
	return chain.Locate(query)
}

func (inst *ComputeBasedLocator) tryLocateAsEnv(query *env.Query) error {
	want := query.Want
	app := want.App
	name := want.Env
	prefix := strings.ToUpper(app) + "_"
	if strings.HasPrefix(name, prefix) {
		key := name[len(prefix):]
		return inst.handleWithKey(query, envRegularPrefix+key)
	}
	return fmt.Errorf("not a env")
}

func (inst *ComputeBasedLocator) tryLocateAsProperty(query *env.Query) error {
	want := query.Want
	app := want.App
	name := want.Property
	prefix := strings.ToLower(app) + "."
	if strings.HasPrefix(name, prefix) {
		key := name[len(prefix):]
		// 把 'property-name' 转换为 'env-name'
		key = strings.ToUpper(key)
		key = strings.ReplaceAll(key, ".", "_")
		return inst.handleWithKey(query, envRegularPrefix+key)
	}
	return fmt.Errorf("not a property")
}

func (inst *ComputeBasedLocator) handleWithKey(query *env.Query, key string) error {
	table := inst.handlers
	if table == nil {
		table = make(map[string]func(q *env.Query) error)
		inst.initKeyHandlers(table)
		inst.handlers = table
	}
	h := table[key]
	if h == nil {
		return fmt.Errorf("no handler for key [%s]", key)
	}
	return h(query)
}

func (inst *ComputeBasedLocator) initKeyHandlers(table map[string]func(q *env.Query) error) {

	lc := &locationComputer{fs: inst.FS}

	table[envUserHome] = lc.locateUserHome
	table[envAppData] = lc.locateAppData
	table[envAppHome] = lc.locateAppHome
	table[envAppConfig] = lc.locateAppConfig
	table[envAppSetting] = lc.locateAppSettings
	table[envUserAppHome] = lc.locateUserAppHome
	table[envShareSystemHome] = lc.locateShareSystemHome
}

////////////////////////////////////////////////////////////////////////////////

type locationComputer struct {
	fs afs.FS
}

func (inst *locationComputer) getUserHome() (afs.Path, error) {
	u, err := user.Current()
	if err != nil {
		return nil, err
	}
	dir := inst.fs.NewPath(u.HomeDir)
	return dir, nil
}

func (inst *locationComputer) getSystemRoot() (afs.Path, error) {

	if runtime.GOOS != "windows" {
		path := inst.fs.NewPath("/")
		return path, nil
	}

	// for windows
	p, err := inst.getUserHome()
	if err != nil {
		return nil, err
	}
	root := p
	for ; p != nil; p = p.GetParent() {
		name := p.GetName()
		if name != "" {
			root = p
		}
	}
	if root == nil {
		return nil, fmt.Errorf("cannot find system root")
	}
	return root, nil
}

func (inst *locationComputer) getUserAppHome(query *env.Query) (afs.Path, error) {
	userHomeDir, err := inst.getUserHome()
	if err != nil {
		return nil, err
	}
	dotAppName := "." + strings.ToLower(query.Want.App)
	dir := userHomeDir.GetChild(dotAppName)
	return dir, nil
}

func (inst *locationComputer) locateUserHome(query *env.Query) error {
	dir, err := inst.getUserHome()
	if err != nil {
		return err
	}
	query.Have.Path = dir
	return nil
}

func (inst *locationComputer) locateAppConfig(query *env.Query) error {
	userAppHome, err := inst.getUserAppHome(query)
	if err != nil {
		return err
	}
	query.Have.Path = userAppHome.GetChild("etc")
	return nil
}

func (inst *locationComputer) locateAppSettings(query *env.Query) error {
	userAppHome, err := inst.getUserAppHome(query)
	if err != nil {
		return err
	}
	query.Have.Path = userAppHome.GetChild("settings")
	return nil
}

func (inst *locationComputer) locateAppData(query *env.Query) error {
	userAppHome, err := inst.getUserAppHome(query)
	if err != nil {
		return err
	}
	query.Have.Path = userAppHome.GetChild("data")
	return nil
}

func (inst *locationComputer) locateUserAppHome(query *env.Query) error {
	userAppHome, err := inst.getUserAppHome(query)
	if err != nil {
		return err
	}
	query.Have.Path = userAppHome
	return nil
}

func (inst *locationComputer) locateAppHome(query *env.Query) error {

	exepath := ""
	for idx, item := range os.Args {
		if idx == 0 {
			exepath = item
			break
		}
	}

	if exepath == "" {
		return fmt.Errorf("cannot locate path of executable")
	}

	exefile := inst.fs.NewPath(exepath)
	if !exefile.IsFile() {
		return fmt.Errorf("no executable file at path [%s]", exepath)
	}

	binDir := exefile.GetParent()
	appHomeDir := binDir
	if binDir.GetName() == "bin" {
		appHomeDir = binDir.GetParent()
	}

	if appHomeDir == nil {
		return fmt.Errorf("cannot locate path of app home dir")
	}
	if !appHomeDir.IsDirectory() {
		return fmt.Errorf("cannot locate path of app home dir")
	}

	query.Have.Path = appHomeDir
	return nil
}

func (inst *locationComputer) locateShareSystemHome(query *env.Query) error {
	root, err := inst.getSystemRoot()
	if err != nil {
		return err
	}
	query.Have.Path = root
	return nil
}
