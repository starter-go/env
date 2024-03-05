package env

import "github.com/starter-go/afs"

// Service ... 对外提供环境信息的服务
type Service interface {
	GetPath(c *Conditions) afs.Path
}
