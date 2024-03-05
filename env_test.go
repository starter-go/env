package env

import (
	"testing"

	"github.com/starter-go/vlog"
)

func TestEnv(t *testing.T) {

	userlist := []User{UserCurrent, UserShare}
	scopelist := []Scope{ScopeApp, ScopeSystem}
	catelist := []Category{CategoryHome, CategoryCode, CategoryData, CategorySetting, CategoryConfig}

	conlist := make([]*Conditions, 0)

	for _, u := range userlist {
		for _, s := range scopelist {
			for _, c := range catelist {
				cond := &Conditions{
					User:     u,
					Scope:    s,
					Category: c,
				}
				conlist = append(conlist, cond)
			}
		}
	}

	vlog.Info("")

	// env
	for _, con := range conlist {
		name := GetEnvName(con)
		vlog.Info("Env Name: %s", name)
	}

	vlog.Info("")

	// property
	for _, con := range conlist {
		name := GetPropertyName(con)
		vlog.Info("Property Name: %s", name)
	}

	vlog.Info("")
}
