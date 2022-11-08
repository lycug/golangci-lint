package golinters

import (
	"github.com/golangci/golangci-lint/pkg/config"
	"github.com/golangci/golangci-lint/pkg/golinters/goanalysis"
	filelenanalyzer "github.com/lycug/filelen/pkg/analyzer"
	"golang.org/x/tools/go/analysis"
)

const filelenName = "filelen"

func NewFileLen(settings *config.FilelenSettings) *goanalysis.Linter {
	analyzer := &analysis.Analyzer{
		Name: filelenName,
		Doc:  goanalysis.TheOnlyanalyzerDoc,
		Run: func(pass *analysis.Pass) (interface{}, error) {
			return filelenanalyzer.Newrun(pass, settings.MaxLineNum)
		},
	}

	return goanalysis.NewLinter(
		funlenName,
		"check file max line number",
		[]*analysis.Analyzer{analyzer},
		nil,
	).WithLoadMode(goanalysis.LoadModeSyntax)
}
