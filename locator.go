package env

import (
	"github.com/starter-go/afs"
)

// Location 表示文件系统中的某个位置
type Location struct {
	Conditions   Conditions
	PriorityName string
	EnvName      string
}

// Locator ...
type Locator interface {
	Locate(want *Location, chain LocatorChain) (afs.Path, error)
}

// LocatorChain ...
type LocatorChain interface {
	Locate(want *Location) (afs.Path, error)
}

// LocatorRegistration ...
type LocatorRegistration struct {
	Locator  Locator
	Enabled  bool
	Priority int // 优先级越高，越靠近处理链条前端

}

// LocatorRegistry ...
type LocatorRegistry interface {
	Registrations() []*LocatorRegistration
}
