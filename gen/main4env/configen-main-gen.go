package main4env

import "github.com/starter-go/application"

func nop(a ... any) {    
}

func registerComponents(cr application.ComponentRegistry) error {
    ac:=&autoRegistrar{}
    ac.init(cr)
    return ac.addAll()
}

type comFactory interface {
    register(cr application.ComponentRegistry) error
}

type autoRegistrar struct {
    cr application.ComponentRegistry
}

func (inst *autoRegistrar) init(cr application.ComponentRegistry) {
	inst.cr = cr
}

func (inst *autoRegistrar) register(factory comFactory) error {
	return factory.register(inst.cr)
}

func (inst*autoRegistrar) addAll() error {

    
    inst.register(&p07222c0995_code_EnvServiceImpl{})
    inst.register(&p1dc79f38c2_locators_CacheLocator{})
    inst.register(&p1dc79f38c2_locators_ComputeBasedLocator{})
    inst.register(&p1dc79f38c2_locators_DefaultLocator{})
    inst.register(&p1dc79f38c2_locators_EnvBasedLocator{})
    inst.register(&p1dc79f38c2_locators_ExampleLocator{})
    inst.register(&p1dc79f38c2_locators_PrepareLocator{})
    inst.register(&p1dc79f38c2_locators_PropertyBasedLocator{})


    return nil
}
