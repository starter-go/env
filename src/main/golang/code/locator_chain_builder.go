package code

import (
	"fmt"
	"sort"

	"github.com/starter-go/env"
)

type resolverChainBuilder struct {
	items []*env.ResolverRegistration
}

func (inst *resolverChainBuilder) add(list ...*env.ResolverRegistration) {
	for _, item := range list {
		if inst.acceptItem(item) {
			inst.items = append(inst.items, item)
		}
	}
}

func (inst *resolverChainBuilder) acceptItem(item *env.ResolverRegistration) bool {
	if item == nil {
		return false
	}
	if item.Resolver == nil {
		return false
	}
	if !item.Enabled {
		return false
	}
	return true
}

func (inst *resolverChainBuilder) create() env.ResolverChain {
	inst.sortItems()
	var chain env.ResolverChain
	chain = new(resolverChainEnding)
	list := inst.items
	for _, reg := range list {
		node := &resolverChainNode{
			next:     chain,
			Resolver: reg.Resolver,
		}
		chain = node
	}
	return chain
}

func (inst *resolverChainBuilder) sortItems() {
	sort.Sort(inst)
}

func (inst *resolverChainBuilder) Swap(i1, i2 int) {
	list := inst.items
	list[i1], list[i2] = list[i2], list[i1]
}

func (inst *resolverChainBuilder) Less(i1, i2 int) bool {
	a := inst.items[i1]
	b := inst.items[i2]
	return a.Priority < b.Priority
}

func (inst *resolverChainBuilder) Len() int {
	return len(inst.items)
}

////////////////////////////////////////////////////////////////////////////////

type resolverChainNode struct {
	next     env.ResolverChain
	Resolver env.Resolver
}

func (inst *resolverChainNode) _impl() env.ResolverChain {
	return inst
}

func (inst *resolverChainNode) Resolve(q *env.Query) error {
	return inst.Resolver.Resolve(q, inst.next)
}

////////////////////////////////////////////////////////////////////////////////

type resolverChainEnding struct{}

func (inst *resolverChainEnding) _impl() env.ResolverChain {
	return inst
}

func (inst *resolverChainEnding) Resolve(q *env.Query) error {
	return fmt.Errorf("cannot handle env.Query")
}
