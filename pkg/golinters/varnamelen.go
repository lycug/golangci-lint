package golinters

import (
	"strconv"
	"strings"

	"github.com/blizzy78/varnamelen"
	"golang.org/x/tools/go/analysis"

	"github.com/lycug/golangci-lint/pkg/config"
	"github.com/lycug/golangci-lint/pkg/golinters/goanalysis"
)

func NewVarnamelen(settings *config.VarnamelenSettings) *goanalysis.Linter {
	analyzer := varnamelen.NewAnalyzer()
	cfg := map[string]map[string]interface{}{}

	if settings != nil {
		vnlCfg := map[string]interface{}{
			"checkReceiver":      strconv.FormatBool(settings.CheckReceiver),
			"checkReturn":        strconv.FormatBool(settings.CheckReturn),
			"checkTypeParam":     strconv.FormatBool(settings.CheckTypeParam),
			"ignoreNames":        strings.Join(settings.IgnoreNames, ","),
			"ignoreTypeAssertOk": strconv.FormatBool(settings.IgnoreTypeAssertOk),
			"ignoreMapIndexOk":   strconv.FormatBool(settings.IgnoreMapIndexOk),
			"ignoreChanRecvOk":   strconv.FormatBool(settings.IgnoreChanRecvOk),
			"ignoreDecls":        strings.Join(settings.IgnoreDecls, ","),
		}

		if settings.MaxDistance > 0 {
			vnlCfg["maxDistance"] = strconv.Itoa(settings.MaxDistance)
		}
		if settings.MinNameLength > 0 {
			vnlCfg["minNameLength"] = strconv.Itoa(settings.MinNameLength)
		}

		cfg[analyzer.Name] = vnlCfg
	}

	return goanalysis.NewLinter(
		analyzer.Name,
		"checks that the length of a variable's name matches its scope",
		[]*analysis.Analyzer{analyzer},
		cfg,
	).WithLoadMode(goanalysis.LoadModeTypesInfo)
}
