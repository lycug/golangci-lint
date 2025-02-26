package commands

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/lycug/golangci-lint/pkg/lint/linter"
)

func (e *Executor) initLinters() {
	e.lintersCmd = &cobra.Command{
		Use:               "linters",
		Short:             "List current linters configuration",
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE:              e.executeLinters,
	}
	e.rootCmd.AddCommand(e.lintersCmd)
	e.initRunConfiguration(e.lintersCmd)
}

// executeLinters runs the 'linters' CLI command, which displays the supported linters.
func (e *Executor) executeLinters(_ *cobra.Command, _ []string) error {
	enabledLintersMap, err := e.EnabledLintersSet.GetEnabledLintersMap()
	if err != nil {
		return fmt.Errorf("can't get enabled linters: %w", err)
	}

	color.Green("Enabled by your configuration linters:\n")
	enabledLinters := make([]*linter.Config, 0, len(enabledLintersMap))
	for _, linter := range enabledLintersMap {
		enabledLinters = append(enabledLinters, linter)
	}
	printLinterConfigs(enabledLinters)

	var disabledLCs []*linter.Config
	for _, lc := range e.DBManager.GetAllSupportedLinterConfigs() {
		if enabledLintersMap[lc.Name()] == nil {
			disabledLCs = append(disabledLCs, lc)
		}
	}

	color.Red("\nDisabled by your configuration linters:\n")
	printLinterConfigs(disabledLCs)

	return nil
}
