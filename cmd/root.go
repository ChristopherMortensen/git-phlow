package cmd

import (
	"fmt"
	"os"

	"github.com/praqma/git-phlow/options"
	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "git-phlow",
	Short: "git extension for phlow workflow",
	Long: `
git-phlow is a git extension that provides an extra set of commands, enabling you to:
workon, deliver and create new tasks.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if options.GlobalFlagVersion != false {
			fmt.Println("phlow specification")
			fmt.Println("sha \t\t: " + options.Sha1)
			fmt.Println("version \t: " + options.Version)
			fmt.Println("date \t\t: " + options.Date)
			os.Exit(0)
		}

		fmt.Println(cmd.UsageString())
	},
}

//Execute ...
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	RootCmd.PersistentFlags()
	RootCmd.Flags().BoolVar(&options.GlobalFlagVersion, "version", false, "current version")

	RootCmd.PersistentFlags().BoolVarP(&options.GlobalFlagVerbose, "verbose", "v", false, "verbose printers")
}
