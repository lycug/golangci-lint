package golinters

import (
	"github.com/kkHAIKE/contextcheck"
	"golang.org/x/tools/go/analysis"

	"github.com/lycug/golangci-lint/pkg/golinters/goanalysis"
	"github.com/lycug/golangci-lint/pkg/lint/linter"
)

func NewContextCheck() *goanalysis.Linter {
	analyzer := contextcheck.NewAnalyzer(contextcheck.Configuration{})

	return goanalysis.NewLinter(
		analyzer.Name,
		analyzer.Doc,
		[]*analysis.Analyzer{analyzer},
		nil,
	).WithContextSetter(func(lintCtx *linter.Context) {
		analyzer.Run = contextcheck.NewRun(lintCtx.Packages, false)
	}).WithLoadMode(goanalysis.LoadModeTypesInfo)
}
