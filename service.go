package env

import (
	"github.com/starter-go/afs"
	"github.com/starter-go/application/environment"
	"github.com/starter-go/application/properties"
)

// Service ... 对外提供环境信息的服务
type Service interface {
	ForApp(appName string) Context
}

// Context ...
type Context interface {
	AppName() string

	Query(c *Want) (*Have, error)

	GetPath(c *Want) afs.Path

	GetEnvironment() environment.Table

	GetProperties() properties.Table
}
