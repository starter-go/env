package code

import (
	"fmt"
	"sort"

	"github.com/starter-go/env"
)

type locatorChainBuilder struct {
	items []*env.LocatorRegistration
}

func (inst *locatorChainBuilder) add(list ...*env.LocatorRegistration) {
	for _, item := range list {
		if inst.acceptItem(item) {
			inst.items = append(inst.items, item)
		}
	}
}

func (inst *locatorChainBuilder) acceptItem(item *env.LocatorRegistration) bool {
	if item == nil {
		return false
	}
	if item.Locator == nil {
		return false
	}
	if !item.Enabled {
		return false
	}
	return true
}

func (inst *locatorChainBuilder) create() env.LocatorChain {
	inst.sortItems()
	var chain env.LocatorChain
	chain = new(locatorChainEnding)
	list := inst.items
	for _, reg := range list {
		node := &locatorChainNode{
			next:    chain,
			locator: reg.Locator,
		}
		chain = node
	}
	return chain
}

func (inst *locatorChainBuilder) sortItems() {
	sort.Sort(inst)
}

func (inst *locatorChainBuilder) Swap(i1, i2 int) {
	list := inst.items
	list[i1], list[i2] = list[i2], list[i1]
}

func (inst *locatorChainBuilder) Less(i1, i2 int) bool {
	a := inst.items[i1]
	b := inst.items[i2]
	return a.Priority < b.Priority
}

func (inst *locatorChainBuilder) Len() int {
	return len(inst.items)
}

////////////////////////////////////////////////////////////////////////////////

type locatorChainNode struct {
	next    env.LocatorChain
	locator env.Locator
}

func (inst *locatorChainNode) _impl() env.LocatorChain {
	return inst
}

func (inst *locatorChainNode) Locate(q *env.Query) error {
	return inst.locator.Locate(q, inst.next)
}

////////////////////////////////////////////////////////////////////////////////

type locatorChainEnding struct{}

func (inst *locatorChainEnding) _impl() env.LocatorChain {
	return inst
}

func (inst *locatorChainEnding) Locate(q *env.Query) error {
	return fmt.Errorf("cannot handle env.Query")
}
