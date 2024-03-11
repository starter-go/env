package main4env
import (
    p0d2a11d16 "github.com/starter-go/afs"
    p22044dd12 "github.com/starter-go/env"
    p07222c099 "github.com/starter-go/env/src/main/golang/code"
    p1dc79f38c "github.com/starter-go/env/src/main/golang/code/locators"
     "github.com/starter-go/application"
)

// type p07222c099.EnvServiceImpl in package:github.com/starter-go/env/src/main/golang/code
//
// id:com-07222c099572b449-code-EnvServiceImpl
// class:
// alias:alias-22044dd124362ad457af360893e95ad0-Service
// scope:singleton
//
type p07222c0995_code_EnvServiceImpl struct {
}

func (inst* p07222c0995_code_EnvServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-07222c099572b449-code-EnvServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-22044dd124362ad457af360893e95ad0-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p07222c0995_code_EnvServiceImpl) new() any {
    return &p07222c099.EnvServiceImpl{}
}

func (inst* p07222c0995_code_EnvServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p07222c099.EnvServiceImpl)
	nop(ie, com)

	
    com.Locators = inst.getLocators(ie)


    return nil
}


func (inst*p07222c0995_code_EnvServiceImpl) getLocators(ie application.InjectionExt)[]p22044dd12.LocatorRegistry{
    dst := make([]p22044dd12.LocatorRegistry, 0)
    src := ie.ListComponents(".class-22044dd124362ad457af360893e95ad0-LocatorRegistry")
    for _, item1 := range src {
        item2 := item1.(p22044dd12.LocatorRegistry)
        dst = append(dst, item2)
    }
    return dst
}



// type p1dc79f38c.CacheLocator in package:github.com/starter-go/env/src/main/golang/code/locators
//
// id:com-1dc79f38c2e0b3af-locators-CacheLocator
// class:class-22044dd124362ad457af360893e95ad0-LocatorRegistry
// alias:
// scope:singleton
//
type p1dc79f38c2_locators_CacheLocator struct {
}

func (inst* p1dc79f38c2_locators_CacheLocator) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-1dc79f38c2e0b3af-locators-CacheLocator"
	r.Classes = "class-22044dd124362ad457af360893e95ad0-LocatorRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p1dc79f38c2_locators_CacheLocator) new() any {
    return &p1dc79f38c.CacheLocator{}
}

func (inst* p1dc79f38c2_locators_CacheLocator) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p1dc79f38c.CacheLocator)
	nop(ie, com)

	


    return nil
}



// type p1dc79f38c.ComputeBasedLocator in package:github.com/starter-go/env/src/main/golang/code/locators
//
// id:com-1dc79f38c2e0b3af-locators-ComputeBasedLocator
// class:class-22044dd124362ad457af360893e95ad0-LocatorRegistry
// alias:
// scope:singleton
//
type p1dc79f38c2_locators_ComputeBasedLocator struct {
}

func (inst* p1dc79f38c2_locators_ComputeBasedLocator) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-1dc79f38c2e0b3af-locators-ComputeBasedLocator"
	r.Classes = "class-22044dd124362ad457af360893e95ad0-LocatorRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p1dc79f38c2_locators_ComputeBasedLocator) new() any {
    return &p1dc79f38c.ComputeBasedLocator{}
}

func (inst* p1dc79f38c2_locators_ComputeBasedLocator) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p1dc79f38c.ComputeBasedLocator)
	nop(ie, com)

	
    com.FS = inst.getFS(ie)


    return nil
}


func (inst*p1dc79f38c2_locators_ComputeBasedLocator) getFS(ie application.InjectionExt)p0d2a11d16.FS{
    return ie.GetComponent("#alias-0d2a11d163e349503a64168a1cdf48a2-FS").(p0d2a11d16.FS)
}



// type p1dc79f38c.DefaultLocator in package:github.com/starter-go/env/src/main/golang/code/locators
//
// id:com-1dc79f38c2e0b3af-locators-DefaultLocator
// class:class-22044dd124362ad457af360893e95ad0-LocatorRegistry
// alias:
// scope:singleton
//
type p1dc79f38c2_locators_DefaultLocator struct {
}

func (inst* p1dc79f38c2_locators_DefaultLocator) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-1dc79f38c2e0b3af-locators-DefaultLocator"
	r.Classes = "class-22044dd124362ad457af360893e95ad0-LocatorRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p1dc79f38c2_locators_DefaultLocator) new() any {
    return &p1dc79f38c.DefaultLocator{}
}

func (inst* p1dc79f38c2_locators_DefaultLocator) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p1dc79f38c.DefaultLocator)
	nop(ie, com)

	
    com.FS = inst.getFS(ie)


    return nil
}


func (inst*p1dc79f38c2_locators_DefaultLocator) getFS(ie application.InjectionExt)p0d2a11d16.FS{
    return ie.GetComponent("#alias-0d2a11d163e349503a64168a1cdf48a2-FS").(p0d2a11d16.FS)
}



// type p1dc79f38c.EnvBasedLocator in package:github.com/starter-go/env/src/main/golang/code/locators
//
// id:com-1dc79f38c2e0b3af-locators-EnvBasedLocator
// class:class-22044dd124362ad457af360893e95ad0-LocatorRegistry
// alias:
// scope:singleton
//
type p1dc79f38c2_locators_EnvBasedLocator struct {
}

func (inst* p1dc79f38c2_locators_EnvBasedLocator) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-1dc79f38c2e0b3af-locators-EnvBasedLocator"
	r.Classes = "class-22044dd124362ad457af360893e95ad0-LocatorRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p1dc79f38c2_locators_EnvBasedLocator) new() any {
    return &p1dc79f38c.EnvBasedLocator{}
}

func (inst* p1dc79f38c2_locators_EnvBasedLocator) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p1dc79f38c.EnvBasedLocator)
	nop(ie, com)

	


    return nil
}



// type p1dc79f38c.ExampleLocator in package:github.com/starter-go/env/src/main/golang/code/locators
//
// id:com-1dc79f38c2e0b3af-locators-ExampleLocator
// class:class-22044dd124362ad457af360893e95ad0-LocatorRegistry
// alias:
// scope:singleton
//
type p1dc79f38c2_locators_ExampleLocator struct {
}

func (inst* p1dc79f38c2_locators_ExampleLocator) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-1dc79f38c2e0b3af-locators-ExampleLocator"
	r.Classes = "class-22044dd124362ad457af360893e95ad0-LocatorRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p1dc79f38c2_locators_ExampleLocator) new() any {
    return &p1dc79f38c.ExampleLocator{}
}

func (inst* p1dc79f38c2_locators_ExampleLocator) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p1dc79f38c.ExampleLocator)
	nop(ie, com)

	


    return nil
}



// type p1dc79f38c.PrepareLocator in package:github.com/starter-go/env/src/main/golang/code/locators
//
// id:com-1dc79f38c2e0b3af-locators-PrepareLocator
// class:class-22044dd124362ad457af360893e95ad0-LocatorRegistry
// alias:
// scope:singleton
//
type p1dc79f38c2_locators_PrepareLocator struct {
}

func (inst* p1dc79f38c2_locators_PrepareLocator) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-1dc79f38c2e0b3af-locators-PrepareLocator"
	r.Classes = "class-22044dd124362ad457af360893e95ad0-LocatorRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p1dc79f38c2_locators_PrepareLocator) new() any {
    return &p1dc79f38c.PrepareLocator{}
}

func (inst* p1dc79f38c2_locators_PrepareLocator) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p1dc79f38c.PrepareLocator)
	nop(ie, com)

	
    com.SimpleAppName = inst.getSimpleAppName(ie)


    return nil
}


func (inst*p1dc79f38c2_locators_PrepareLocator) getSimpleAppName(ie application.InjectionExt)string{
    return ie.GetString("${application.simple-name}")
}



// type p1dc79f38c.PropertyBasedLocator in package:github.com/starter-go/env/src/main/golang/code/locators
//
// id:com-1dc79f38c2e0b3af-locators-PropertyBasedLocator
// class:class-22044dd124362ad457af360893e95ad0-LocatorRegistry
// alias:
// scope:singleton
//
type p1dc79f38c2_locators_PropertyBasedLocator struct {
}

func (inst* p1dc79f38c2_locators_PropertyBasedLocator) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-1dc79f38c2e0b3af-locators-PropertyBasedLocator"
	r.Classes = "class-22044dd124362ad457af360893e95ad0-LocatorRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p1dc79f38c2_locators_PropertyBasedLocator) new() any {
    return &p1dc79f38c.PropertyBasedLocator{}
}

func (inst* p1dc79f38c2_locators_PropertyBasedLocator) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p1dc79f38c.PropertyBasedLocator)
	nop(ie, com)

	


    return nil
}


