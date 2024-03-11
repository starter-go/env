package code

import (
	"sort"

	"github.com/starter-go/env"
	"github.com/starter-go/units"
	"github.com/starter-go/vlog"
)

// DemoUnit ... 单元测试示例
type DemoUnit struct {

	//starter:component
	_as func(units.Units) //starter:as(".")

	Service env.Service //starter:inject("#")

}

func (inst *DemoUnit) _impl() units.Units { return inst }

// Units ...
func (inst *DemoUnit) Units(list []*units.Registration) []*units.Registration {

	list = append(list, &units.Registration{
		Name:     "test-log-env",
		Enabled:  true,
		Priority: 0,
		Test:     inst.test1,
	})

	return list
}

func (inst *DemoUnit) test1() error {

	ctx := inst.Service.ForApp("DEMO")
	allEnv := ctx.GetEnvironment().Export(nil)
	allPro := ctx.GetProperties().Export(nil)

	inst.logTable("env:", allEnv)
	inst.logTable("properties:", allPro)

	return nil
}

func (inst *DemoUnit) logTable(tag string, table map[string]string) {
	vlog.Info(tag)
	keys := make([]string, 0)
	for key := range table {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		value := table[key]
		vlog.Info("  %s = %s", key, value)
	}
}
