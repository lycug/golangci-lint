package printers

import (
	"context"

	"github.com/lycug/golangci-lint/pkg/result"
)

type Printer interface {
	Print(ctx context.Context, issues []result.Issue) error
}
