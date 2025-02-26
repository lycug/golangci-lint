package golinters

import (
	"github.com/sashamelentyev/interfacebloat/pkg/analyzer"
	"golang.org/x/tools/go/analysis"

	"github.com/lycug/golangci-lint/pkg/config"
	"github.com/lycug/golangci-lint/pkg/golinters/goanalysis"
)

func NewInterfaceBloat(settings *config.InterfaceBloatSettings) *goanalysis.Linter {
	a := analyzer.New()

	var cfg map[string]map[string]interface{}
	if settings != nil {
		cfg = map[string]map[string]interface{}{
			a.Name: {
				analyzer.InterfaceMaxMethodsFlag: settings.Max,
			},
		}
	}

	return goanalysis.NewLinter(
		a.Name,
		a.Doc,
		[]*analysis.Analyzer{a},
		cfg,
	).WithLoadMode(goanalysis.LoadModeSyntax)
}
