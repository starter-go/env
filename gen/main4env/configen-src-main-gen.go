package main4env
import (
    p0d2a11d16 "github.com/starter-go/afs"
    p0ef6f2938 "github.com/starter-go/application"
    p22044dd12 "github.com/starter-go/env"
    p07222c099 "github.com/starter-go/env/src/main/golang/code"
    p3a6fd730c "github.com/starter-go/env/src/main/golang/code/resolvers"
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

	
    com.Resolvers = inst.getResolvers(ie)


    return nil
}


func (inst*p07222c0995_code_EnvServiceImpl) getResolvers(ie application.InjectionExt)[]p22044dd12.ResolverRegistry{
    dst := make([]p22044dd12.ResolverRegistry, 0)
    src := ie.ListComponents(".class-22044dd124362ad457af360893e95ad0-ResolverRegistry")
    for _, item1 := range src {
        item2 := item1.(p22044dd12.ResolverRegistry)
        dst = append(dst, item2)
    }
    return dst
}



// type p3a6fd730c.CacheResolver in package:github.com/starter-go/env/src/main/golang/code/resolvers
//
// id:com-3a6fd730c75a4006-resolvers-CacheResolver
// class:class-22044dd124362ad457af360893e95ad0-ResolverRegistry
// alias:
// scope:singleton
//
type p3a6fd730c7_resolvers_CacheResolver struct {
}

func (inst* p3a6fd730c7_resolvers_CacheResolver) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-3a6fd730c75a4006-resolvers-CacheResolver"
	r.Classes = "class-22044dd124362ad457af360893e95ad0-ResolverRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p3a6fd730c7_resolvers_CacheResolver) new() any {
    return &p3a6fd730c.CacheResolver{}
}

func (inst* p3a6fd730c7_resolvers_CacheResolver) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p3a6fd730c.CacheResolver)
	nop(ie, com)

	


    return nil
}



// type p3a6fd730c.ComputeBasedResolver in package:github.com/starter-go/env/src/main/golang/code/resolvers
//
// id:com-3a6fd730c75a4006-resolvers-ComputeBasedResolver
// class:class-22044dd124362ad457af360893e95ad0-ResolverRegistry
// alias:
// scope:singleton
//
type p3a6fd730c7_resolvers_ComputeBasedResolver struct {
}

func (inst* p3a6fd730c7_resolvers_ComputeBasedResolver) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-3a6fd730c75a4006-resolvers-ComputeBasedResolver"
	r.Classes = "class-22044dd124362ad457af360893e95ad0-ResolverRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p3a6fd730c7_resolvers_ComputeBasedResolver) new() any {
    return &p3a6fd730c.ComputeBasedResolver{}
}

func (inst* p3a6fd730c7_resolvers_ComputeBasedResolver) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p3a6fd730c.ComputeBasedResolver)
	nop(ie, com)

	
    com.FS = inst.getFS(ie)


    return nil
}


func (inst*p3a6fd730c7_resolvers_ComputeBasedResolver) getFS(ie application.InjectionExt)p0d2a11d16.FS{
    return ie.GetComponent("#alias-0d2a11d163e349503a64168a1cdf48a2-FS").(p0d2a11d16.FS)
}



// type p3a6fd730c.DefaultResolver in package:github.com/starter-go/env/src/main/golang/code/resolvers
//
// id:com-3a6fd730c75a4006-resolvers-DefaultResolver
// class:class-22044dd124362ad457af360893e95ad0-ResolverRegistry
// alias:
// scope:singleton
//
type p3a6fd730c7_resolvers_DefaultResolver struct {
}

func (inst* p3a6fd730c7_resolvers_DefaultResolver) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-3a6fd730c75a4006-resolvers-DefaultResolver"
	r.Classes = "class-22044dd124362ad457af360893e95ad0-ResolverRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p3a6fd730c7_resolvers_DefaultResolver) new() any {
    return &p3a6fd730c.DefaultResolver{}
}

func (inst* p3a6fd730c7_resolvers_DefaultResolver) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p3a6fd730c.DefaultResolver)
	nop(ie, com)

	
    com.FS = inst.getFS(ie)


    return nil
}


func (inst*p3a6fd730c7_resolvers_DefaultResolver) getFS(ie application.InjectionExt)p0d2a11d16.FS{
    return ie.GetComponent("#alias-0d2a11d163e349503a64168a1cdf48a2-FS").(p0d2a11d16.FS)
}



// type p3a6fd730c.EnvBasedResolver in package:github.com/starter-go/env/src/main/golang/code/resolvers
//
// id:com-3a6fd730c75a4006-resolvers-EnvBasedResolver
// class:class-22044dd124362ad457af360893e95ad0-ResolverRegistry
// alias:
// scope:singleton
//
type p3a6fd730c7_resolvers_EnvBasedResolver struct {
}

func (inst* p3a6fd730c7_resolvers_EnvBasedResolver) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-3a6fd730c75a4006-resolvers-EnvBasedResolver"
	r.Classes = "class-22044dd124362ad457af360893e95ad0-ResolverRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p3a6fd730c7_resolvers_EnvBasedResolver) new() any {
    return &p3a6fd730c.EnvBasedResolver{}
}

func (inst* p3a6fd730c7_resolvers_EnvBasedResolver) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p3a6fd730c.EnvBasedResolver)
	nop(ie, com)

	
    com.AppContext = inst.getAppContext(ie)


    return nil
}


func (inst*p3a6fd730c7_resolvers_EnvBasedResolver) getAppContext(ie application.InjectionExt)p0ef6f2938.Context{
    return ie.GetContext()
}



// type p3a6fd730c.ExampleResolver in package:github.com/starter-go/env/src/main/golang/code/resolvers
//
// id:com-3a6fd730c75a4006-resolvers-ExampleResolver
// class:class-22044dd124362ad457af360893e95ad0-ResolverRegistry
// alias:
// scope:singleton
//
type p3a6fd730c7_resolvers_ExampleResolver struct {
}

func (inst* p3a6fd730c7_resolvers_ExampleResolver) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-3a6fd730c75a4006-resolvers-ExampleResolver"
	r.Classes = "class-22044dd124362ad457af360893e95ad0-ResolverRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p3a6fd730c7_resolvers_ExampleResolver) new() any {
    return &p3a6fd730c.ExampleResolver{}
}

func (inst* p3a6fd730c7_resolvers_ExampleResolver) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p3a6fd730c.ExampleResolver)
	nop(ie, com)

	


    return nil
}



// type p3a6fd730c.MutexResolver in package:github.com/starter-go/env/src/main/golang/code/resolvers
//
// id:com-3a6fd730c75a4006-resolvers-MutexResolver
// class:class-22044dd124362ad457af360893e95ad0-ResolverRegistry
// alias:
// scope:singleton
//
type p3a6fd730c7_resolvers_MutexResolver struct {
}

func (inst* p3a6fd730c7_resolvers_MutexResolver) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-3a6fd730c75a4006-resolvers-MutexResolver"
	r.Classes = "class-22044dd124362ad457af360893e95ad0-ResolverRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p3a6fd730c7_resolvers_MutexResolver) new() any {
    return &p3a6fd730c.MutexResolver{}
}

func (inst* p3a6fd730c7_resolvers_MutexResolver) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p3a6fd730c.MutexResolver)
	nop(ie, com)

	


    return nil
}



// type p3a6fd730c.NamesResolver in package:github.com/starter-go/env/src/main/golang/code/resolvers
//
// id:com-3a6fd730c75a4006-resolvers-NamesResolver
// class:class-22044dd124362ad457af360893e95ad0-ResolverRegistry
// alias:
// scope:singleton
//
type p3a6fd730c7_resolvers_NamesResolver struct {
}

func (inst* p3a6fd730c7_resolvers_NamesResolver) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-3a6fd730c75a4006-resolvers-NamesResolver"
	r.Classes = "class-22044dd124362ad457af360893e95ad0-ResolverRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p3a6fd730c7_resolvers_NamesResolver) new() any {
    return &p3a6fd730c.NamesResolver{}
}

func (inst* p3a6fd730c7_resolvers_NamesResolver) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p3a6fd730c.NamesResolver)
	nop(ie, com)

	


    return nil
}



// type p3a6fd730c.PrepareResolver in package:github.com/starter-go/env/src/main/golang/code/resolvers
//
// id:com-3a6fd730c75a4006-resolvers-PrepareResolver
// class:class-22044dd124362ad457af360893e95ad0-ResolverRegistry
// alias:
// scope:singleton
//
type p3a6fd730c7_resolvers_PrepareResolver struct {
}

func (inst* p3a6fd730c7_resolvers_PrepareResolver) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-3a6fd730c75a4006-resolvers-PrepareResolver"
	r.Classes = "class-22044dd124362ad457af360893e95ad0-ResolverRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p3a6fd730c7_resolvers_PrepareResolver) new() any {
    return &p3a6fd730c.PrepareResolver{}
}

func (inst* p3a6fd730c7_resolvers_PrepareResolver) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p3a6fd730c.PrepareResolver)
	nop(ie, com)

	
    com.SimpleAppName = inst.getSimpleAppName(ie)
    com.FS = inst.getFS(ie)


    return nil
}


func (inst*p3a6fd730c7_resolvers_PrepareResolver) getSimpleAppName(ie application.InjectionExt)string{
    return ie.GetString("${application.simple-name}")
}


func (inst*p3a6fd730c7_resolvers_PrepareResolver) getFS(ie application.InjectionExt)p0d2a11d16.FS{
    return ie.GetComponent("#alias-0d2a11d163e349503a64168a1cdf48a2-FS").(p0d2a11d16.FS)
}



// type p3a6fd730c.PropertyBasedResolver in package:github.com/starter-go/env/src/main/golang/code/resolvers
//
// id:com-3a6fd730c75a4006-resolvers-PropertyBasedResolver
// class:class-22044dd124362ad457af360893e95ad0-ResolverRegistry
// alias:
// scope:singleton
//
type p3a6fd730c7_resolvers_PropertyBasedResolver struct {
}

func (inst* p3a6fd730c7_resolvers_PropertyBasedResolver) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-3a6fd730c75a4006-resolvers-PropertyBasedResolver"
	r.Classes = "class-22044dd124362ad457af360893e95ad0-ResolverRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p3a6fd730c7_resolvers_PropertyBasedResolver) new() any {
    return &p3a6fd730c.PropertyBasedResolver{}
}

func (inst* p3a6fd730c7_resolvers_PropertyBasedResolver) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p3a6fd730c.PropertyBasedResolver)
	nop(ie, com)

	
    com.AppContext = inst.getAppContext(ie)


    return nil
}


func (inst*p3a6fd730c7_resolvers_PropertyBasedResolver) getAppContext(ie application.InjectionExt)p0ef6f2938.Context{
    return ie.GetContext()
}


