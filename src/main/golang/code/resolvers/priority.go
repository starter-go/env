package resolvers

// 定义各个定位器的优先级
const (
	PriorityMutex   = 9999
	PriorityPrepare = 9998

	PriorityCache         = 900
	PriorityEnvBased      = 800
	PriorityPropertyBased = 700
	PriorityComputeBased  = 600

	PriorityNames   = 2
	PriorityExample = 1
	PriorityDefault = 0
)
