package locators

// 定义各个定位器的优先级
const (
	PriorityPrepare = 9999

	PriorityCache         = 900
	PriorityEnvBased      = 800
	PriorityPropertyBased = 700
	PriorityComputeBased  = 600

	PriorityExample = 1
	PriorityDefault = 0
)
