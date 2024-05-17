package module

import "github.com/advanced-go/stdlib/core"

const (
	Authority = "github/advanced-go/observation"
	Name      = "observation"
	Version   = "2.2.2"
)

func Into() core.ModuleInfo {
	return core.ModuleInfo{
		Authority: Authority,
		Version:   Version,
		Name:      Name,
	}
}
