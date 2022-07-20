package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "md-proofer",
	Short: "Tool to test Markdown files",
	Long: `Markdown Proofer (md-proofer) is a tool to help test Markdown files.

The main use case envisioned is to test Markdown files before building them 
a static site generator such as Hugo or Jekyll.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
