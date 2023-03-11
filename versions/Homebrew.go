package versions

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os/exec"
)

type homebrew struct{}

var Homebrew homebrew

func (homebrew) Version() {
	cmd := exec.Command("brew", "--version")
	if viper.GetBool("debug") {
		fmt.Println(cmd.String())
	}

	out, err := cmd.Output()
	cobra.CheckErr(err)
	outStr := trimNewline(string(out))

	if viper.GetBool("verbose") {
		fmt.Println("Homebrew version:", outStr)
	} else if viper.GetBool("succinct") {
		versionNumber := outStr[len("Homebrew "):]
		fmt.Println(versionNumber)
	} else {
		fmt.Println(outStr)
	}
}
