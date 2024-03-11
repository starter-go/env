package test4env
import (
    p22044dd12 "github.com/starter-go/env"
    pe4f7c64b7 "github.com/starter-go/env/src/test/golang/code"
     "github.com/starter-go/application"
)

// type pe4f7c64b7.DemoUnit in package:github.com/starter-go/env/src/test/golang/code
//
// id:com-e4f7c64b7c70a76b-code-DemoUnit
// class:class-0dc072ed44b3563882bff4e657a52e62-Units
// alias:
// scope:singleton
//
type pe4f7c64b7c_code_DemoUnit struct {
}

func (inst* pe4f7c64b7c_code_DemoUnit) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-e4f7c64b7c70a76b-code-DemoUnit"
	r.Classes = "class-0dc072ed44b3563882bff4e657a52e62-Units"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pe4f7c64b7c_code_DemoUnit) new() any {
    return &pe4f7c64b7.DemoUnit{}
}

func (inst* pe4f7c64b7c_code_DemoUnit) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pe4f7c64b7.DemoUnit)
	nop(ie, com)

	
    com.Service = inst.getService(ie)


    return nil
}


func (inst*pe4f7c64b7c_code_DemoUnit) getService(ie application.InjectionExt)p22044dd12.Service{
    return ie.GetComponent("#alias-22044dd124362ad457af360893e95ad0-Service").(p22044dd12.Service)
}


