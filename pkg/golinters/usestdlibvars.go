package golinters

import (
	"github.com/sashamelentyev/usestdlibvars/pkg/analyzer"
	"golang.org/x/tools/go/analysis"

	"github.com/lycug/golangci-lint/pkg/config"
	"github.com/lycug/golangci-lint/pkg/golinters/goanalysis"
)

func NewUseStdlibVars(cfg *config.UseStdlibVarsSettings) *goanalysis.Linter {
	a := analyzer.New()

	cfgMap := make(map[string]map[string]interface{})
	if cfg != nil {
		cfgMap[a.Name] = map[string]interface{}{
			analyzer.ConstantKindFlag:       cfg.ConstantKind,
			analyzer.CryptoHashFlag:         cfg.CryptoHash,
			analyzer.HTTPMethodFlag:         cfg.HTTPMethod,
			analyzer.HTTPStatusCodeFlag:     cfg.HTTPStatusCode,
			analyzer.OSDevNullFlag:          cfg.OSDevNull,
			analyzer.RPCDefaultPathFlag:     cfg.DefaultRPCPath,
			analyzer.SQLIsolationLevelFlag:  cfg.SQLIsolationLevel,
			analyzer.SyslogPriorityFlag:     cfg.SyslogPriority,
			analyzer.TimeLayoutFlag:         cfg.TimeLayout,
			analyzer.TimeMonthFlag:          cfg.TimeMonth,
			analyzer.TimeWeekdayFlag:        cfg.TimeWeekday,
			analyzer.TLSSignatureSchemeFlag: cfg.TLSSignatureScheme,
		}
	}

	return goanalysis.NewLinter(
		a.Name,
		a.Doc,
		[]*analysis.Analyzer{a},
		cfgMap,
	).WithLoadMode(goanalysis.LoadModeSyntax)
}
