package Reporter

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	. "github.com/hisaac/sysport/utils"
)

type macOS struct{}

var MacOS macOS

func (os macOS) Info() {
	// TODO: Implement this
}

func (macOS) Version() {
	productNameCmd := exec.Command("sw_vers", "-productName")
	if viper.GetBool("debug") {
		fmt.Println(productNameCmd.String())
	}
	productName, err := productNameCmd.Output()
	cobra.CheckErr(err)
	productNameStr := TrimNewline(string(productName))

	productVersionCmd := exec.Command("sw_vers", "-productVersion")
	if viper.GetBool("debug") {
		fmt.Println(productVersionCmd.String())
	}
	productVersion, err := productVersionCmd.Output()
	cobra.CheckErr(err)
	productVersionStr := TrimNewline(string(productVersion))

	buildVersionCmd := exec.Command("sw_vers", "-buildVersion")
	if viper.GetBool("debug") {
		fmt.Println(buildVersionCmd.String())
	}
	buildVersion, err := buildVersionCmd.Output()
	cobra.CheckErr(err)
	buildVersionStr := TrimNewline(string(buildVersion))

	if viper.GetBool("succinct") {
		fmt.Println(productVersionStr)
		return
	}

	versionString := productNameStr + " " + productVersionStr + " (" + buildVersionStr + ")"

	if viper.GetBool("verbose") {
		fmt.Println("macOS version:", versionString)
	} else {
		fmt.Println(versionString)
	}
}
