package golinters

import (
	"github.com/moricho/tparallel"
	"golang.org/x/tools/go/analysis"

	"github.com/lycug/golangci-lint/pkg/golinters/goanalysis"
)

func NewTparallel() *goanalysis.Linter {
	return goanalysis.NewLinter(
		"tparallel",
		"tparallel detects inappropriate usage of t.Parallel() method in your Go test codes",
		[]*analysis.Analyzer{tparallel.Analyzer},
		nil,
	).WithLoadMode(goanalysis.LoadModeTypesInfo)
}
