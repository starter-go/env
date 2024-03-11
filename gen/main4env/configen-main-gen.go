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
    inst.register(&p3a6fd730c7_resolvers_CacheResolver{})
    inst.register(&p3a6fd730c7_resolvers_ComputeBasedResolver{})
    inst.register(&p3a6fd730c7_resolvers_DefaultResolver{})
    inst.register(&p3a6fd730c7_resolvers_EnvBasedResolver{})
    inst.register(&p3a6fd730c7_resolvers_ExampleResolver{})
    inst.register(&p3a6fd730c7_resolvers_MutexResolver{})
    inst.register(&p3a6fd730c7_resolvers_NamesResolver{})
    inst.register(&p3a6fd730c7_resolvers_PrepareResolver{})
    inst.register(&p3a6fd730c7_resolvers_PropertyBasedResolver{})


    return nil
}
