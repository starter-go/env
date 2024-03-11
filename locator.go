package env

// Query 表示查询文件系统中的某个位置
type Query struct {
	Want *Want
	Have *Have
}

// Locator ...
type Locator interface {
	Locate(q *Query, chain LocatorChain) error
}

// LocatorChain ...
type LocatorChain interface {
	Locate(q *Query) error
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
