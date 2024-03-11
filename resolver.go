package env

// Query 表示查询某个环境变量（或属性）
type Query struct {
	Want *Want
	Have *Have

	// 这个字段用于列举支持的查询名称（可选项）；
	// 如果需要查询此项，用户需要在查询时提供一个空的 map 在这里；
	Names map[string]*Want
}

// Resolver ...
type Resolver interface {
	Resolve(q *Query, chain ResolverChain) error
}

// ResolverChain ...
type ResolverChain interface {
	Resolve(q *Query) error
}

// ResolverRegistration ...
type ResolverRegistration struct {
	Resolver Resolver
	Enabled  bool
	Priority int // 优先级越高，越靠近处理链条前端
}

// ResolverRegistry ...
type ResolverRegistry interface {
	Registrations() []*ResolverRegistration
}
