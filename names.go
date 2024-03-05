package env

import "strings"

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

// Conditions ...
type Conditions struct {
	AppName  string
	Category Category
	Scope    Scope
	User     User
}

// GetEnvName 根据条件取环境变量名称
func GetEnvName(c *Conditions) string {
	str := getVarName(c, "_")
	return strings.ToUpper(str)
}

// GetPropertyName 根据条件取属性名称
func GetPropertyName(c *Conditions) string {
	str := getVarName(c, ".")
	return strings.ToLower(str)
}

func getVarName(c *Conditions, sep string) string {

	app := ""
	user := ""
	scope := ""
	category := ""

	if c != nil {
		app = c.AppName
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
