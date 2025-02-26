package golinters

import (
	"github.com/nishanths/exhaustive"
	"golang.org/x/tools/go/analysis"

	"github.com/lycug/golangci-lint/pkg/config"
	"github.com/lycug/golangci-lint/pkg/golinters/goanalysis"
)

func NewExhaustive(settings *config.ExhaustiveSettings) *goanalysis.Linter {
	a := exhaustive.Analyzer

	var cfg map[string]map[string]interface{}
	if settings != nil {
		cfg = map[string]map[string]interface{}{
			a.Name: {
				exhaustive.CheckFlag:                      settings.Check,
				exhaustive.CheckGeneratedFlag:             settings.CheckGenerated,
				exhaustive.DefaultSignifiesExhaustiveFlag: settings.DefaultSignifiesExhaustive,
				exhaustive.IgnoreEnumMembersFlag:          settings.IgnoreEnumMembers,
				exhaustive.PackageScopeOnlyFlag:           settings.PackageScopeOnly,
				exhaustive.ExplicitExhaustiveMapFlag:      settings.ExplicitExhaustiveMap,
				exhaustive.ExplicitExhaustiveSwitchFlag:   settings.ExplicitExhaustiveSwitch,
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
