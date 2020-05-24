package main

import (
	"github.com/spf13/viper"

	"github.com/felicianotech/md-proofer/cmd"
)

var (
	buildTime string = ""
	gitCommit string = ""
	version   string = "dev-build"
)

func main() {

	viper.SetDefault("buildTime", buildTime)
	viper.SetDefault("gitCommit", gitCommit)
	viper.SetDefault("version", version)

	cmd.Execute()
}
