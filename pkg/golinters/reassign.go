package golinters

import (
	"fmt"
	"strings"

	"github.com/curioswitch/go-reassign"
	"golang.org/x/tools/go/analysis"

	"github.com/lycug/golangci-lint/pkg/config"
	"github.com/lycug/golangci-lint/pkg/golinters/goanalysis"
)

func NewReassign(settings *config.ReassignSettings) *goanalysis.Linter {
	a := reassign.NewAnalyzer()

	var cfg map[string]map[string]interface{}
	if settings != nil && len(settings.Patterns) > 0 {
		cfg = map[string]map[string]interface{}{
			a.Name: {
				reassign.FlagPattern: fmt.Sprintf("^(%s)$", strings.Join(settings.Patterns, "|")),
			},
		}
	}

	return goanalysis.NewLinter(
		a.Name,
		a.Doc,
		[]*analysis.Analyzer{a},
		cfg,
	).WithLoadMode(goanalysis.LoadModeTypesInfo)
}
