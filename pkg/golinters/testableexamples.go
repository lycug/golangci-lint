package golinters

import (
	"github.com/maratori/testableexamples/pkg/testableexamples"
	"golang.org/x/tools/go/analysis"

	"github.com/lycug/golangci-lint/pkg/golinters/goanalysis"
)

func NewTestableexamples() *goanalysis.Linter {
	a := testableexamples.NewAnalyzer()

	return goanalysis.NewLinter(
		a.Name,
		a.Doc,
		[]*analysis.Analyzer{a},
		nil,
	).WithLoadMode(goanalysis.LoadModeSyntax)
}
