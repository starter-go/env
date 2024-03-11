package env

import (
	"github.com/starter-go/application"
	"github.com/starter-go/env"
	"github.com/starter-go/env/gen/main4env"
	"github.com/starter-go/env/gen/test4env"
	"github.com/starter-go/libafs/modules/libafs"
)

// Module  ...
func Module() application.Module {
	mb := env.NewMainModule()
	mb.Components(main4env.ExportComponents)
	return mb.Create()
}

// ModuleForTest ...
func ModuleForTest() application.Module {

	target := Module()

	mb := env.NewTestModule()
	mb.Components(test4env.ExportComponents)
	mb.Depend(target)
	mb.Depend(libafs.Module())

	return mb.Create()
}
