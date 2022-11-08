//golangcitest:args -Egci
//golangcitest:config_path testdata/configs/gci.yml
package testdata

import (
	"fmt"

	"github.com/lycug/golangci-lint/pkg/config" // want "File is not \\`gci\\`-ed with --skip-generated -s standard,prefix\\(github.com/lycug/golangci-lint\\),default"

	"github.com/pkg/errors" // want "File is not \\`gci\\`-ed with --skip-generated -s standard,prefix\\(github.com/lycug/golangci-lint\\),default"
)

func GoimportsLocalTest() {
	fmt.Print("x")
	_ = config.Config{}
	_ = errors.New("")
}
