package golinters

import "github.com/lycug/golangci-lint/pkg/logutils"

// linterLogger must be use only when the context logger is not available.
var linterLogger = logutils.NewStderrLog(logutils.DebugKeyLinter)
