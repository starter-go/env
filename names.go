package env

import (
	"strings"

	"github.com/starter-go/afs"
)

// User ...用户范围[current|share]
type User string

// Scope ...作用域[application|system]
type Scope string

// Category ...内容分类 [data|code|config|setting|...]
type Category string

// 定义内容分类
const (
	CategoryData    Category = "data"
	CategoryCode    Category = "code"
	CategoryConfig  Category = "config"
	CategorySetting Category = "setting"
	CategoryHome    Category = "home"
)

// 定义用户范围
const (
	UserDefault User = ""
	UserCurrent User = "current"
	UserShare   User = "share"
)

// 定义作用域
const (
	ScopeDefault Scope = ""
	ScopeApp     Scope = "application"
	ScopeSystem  Scope = "system"
)

// Want ... 表示查询条件
type Want struct {
	App      string // the simple-app-name
	Env      string // 环境变量名称
	Property string // 属性名称

	Category Category
	Scope    Scope
	User     User
}

// Have ... 表示查询结果
type Have struct {
	Want *Want
	Path afs.Path
}

// ComputeEnvName 根据条件取环境变量名称
func ComputeEnvName(c *Want) string {
	str := computeVarName(c, "_")
	return strings.ToUpper(str)
}

// ComputePropertyName 根据条件取属性名称
func ComputePropertyName(c *Want) string {
	str := computeVarName(c, ".")
	return strings.ToLower(str)
}

// ListPropertyNames 列出可能的属性名称
func ListPropertyNames(app string) []string {
	return listNamesFor(app, ComputePropertyName)
}

// ListEnvNames 列出可能的环境变量名称
func ListEnvNames(app string) []string {
	return listNamesFor(app, ComputeEnvName)
}

func listNamesFor(app string, nameMapper func(c *Want) string) []string {
	if app == "" {
		app = theDefaultAppName
	}
	if app == "" {
		app = "demo"
	}
	userlist := []User{UserCurrent, UserShare, ""}
	scopelist := []Scope{ScopeApp, ScopeSystem, ""}
	catelist := []Category{CategoryHome, CategoryCode, CategoryData, CategorySetting, CategoryConfig}
	wantlist := make([]*Want, 0)
	for _, u := range userlist {
		for _, s := range scopelist {
			for _, c := range catelist {
				cond := &Want{
					User:     u,
					Scope:    s,
					Category: c,
				}
				wantlist = append(wantlist, cond)
			}
		}
	}
	dst := make([]string, 0)
	for _, want := range wantlist {
		want.App = app
		name := nameMapper(want)
		dst = append(dst, name)
	}
	return dst
}

func computeVarName(c *Want, sep string) string {

	app := ""
	user := ""
	scope := ""
	category := ""

	if c != nil {
		app = c.App
		scope = string(c.Scope)
		user = string(c.User)
		category = string(c.Category)
	}

	if app == "" {
		app = "demo"
	}

	b := &strings.Builder{}
	items := []string{app, user, scope, category}
	count := 0

	for _, item := range items {
		if item == "" {
			continue
		}
		if count > 0 {
			b.WriteString(sep)
		}
		b.WriteString(item)
		count++
	}

	return b.String()
}

var theDefaultAppName string

// SetDefaultAppName 设置默认的应用程序名称
// 这个函数是一次性的，一旦设置成功，后续的调用将无效
func SetDefaultAppName(name string) {
	if theDefaultAppName != "" {
		return
	}
	theDefaultAppName = name
}
