package env

import (
	"testing"

	"github.com/starter-go/vlog"
)

func TestEnvNames(t *testing.T) {
	namelist := ListEnvNames("et")
	for _, name := range namelist {
		vlog.Info("Env Name: %s", name)
	}
}

func TestPropertyNames(t *testing.T) {
	namelist := ListPropertyNames("et")
	for _, name := range namelist {
		vlog.Info("Property Name: %s", name)
	}
}
