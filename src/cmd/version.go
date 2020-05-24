package cmd

import (
	"fmt"
	"runtime"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tcnksm/go-latest"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Returns the Semantic Versioning (SemVer) version of Markdown Proofer.",
	Run: func(cmd *cobra.Command, args []string) {
		version := viper.GetString("version")

		if version != "dev-build" {
			githubTag := &latest.GithubTag{
				Owner:      "felicianotech",
				Repository: "md-proofer",
			}

			res, _ := latest.Check(githubTag, version)

			fmt.Printf("Markdown Proofer v%s\n", version)

			if res.Outdated {
				fmt.Printf("An update is available (v%s). Please consider upgrading.", res.Current)
			}
		} else {
			fmt.Printf("Markdown Proofer %s %s/%s BuildDate: %s\n", version, runtime.GOOS, runtime.GOARCH, time.Now().Format(time.RFC3339))
		}

	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
