package Reporter

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	. "github.com/hisaac/sysport/utils"
)

type homebrew struct{}

var Homebrew homebrew

func (h homebrew) Info() {
	// TODO: Implement this
}

func (homebrew) Version() {
	cmd := exec.Command("brew", "--version")
	if viper.GetBool("debug") {
		fmt.Println(cmd.String())
	}

	out, err := cmd.Output()
	cobra.CheckErr(err)
	outStr := TrimNewline(string(out))

	if viper.GetBool("verbose") {
		fmt.Println("Homebrew version:", outStr)
	} else if viper.GetBool("succinct") {
		versionNumber := outStr[len("Homebrew "):]
		fmt.Println(versionNumber)
	} else {
		fmt.Println(outStr)
	}
}
