package golinters

import (
	filelenanalyzer "github.com/lycug/filelen/pkg/analyzer"
	"github.com/lycug/golangci-lint/pkg/config"
	"github.com/lycug/golangci-lint/pkg/golinters/goanalysis"
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
		filelenName,
		"check file max line number",
		[]*analysis.Analyzer{analyzer},
		nil,
	).WithLoadMode(goanalysis.LoadModeSyntax)
}
