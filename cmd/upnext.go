package cmd

import (
	"fmt"

	"github.com/praqma/git-phlow/cmdcheck"
	"github.com/praqma/git-phlow/githandler"
	"github.com/praqma/git-phlow/options"
	"github.com/praqma/git-phlow/phlow"
	"github.com/praqma/git-phlow/plugins"
	"github.com/praqma/git-phlow/ui"
	"github.com/spf13/cobra"
)

// upNextCmd represents the upnext command
var upNextCmd = &cobra.Command{
	Use:   "upnext",
	Short: "get the chronologically next branch",
	Long: fmt.Sprintf(`
%s gets the next branch ready for integration based on the branch creation time.
The branch created first, is the branch thats up next.
`, ui.Bold("upnext")),
	PreRun: func(cmd *cobra.Command, args []string) {
		checks.IsRepository()
	},
	Run: func(cmd *cobra.Command, args []string) {

		defaultBranch, _ := plugins.GitHub.Branch.Default()
		remote := githandler.ConfigBranchRemote(defaultBranch)

		phlow.UpNext(remote)
	},
}

func init() {
	//Adding UpNext as sub-command to agentCmd
	agentCmd.AddCommand(upNextCmd)

	upNextCmd.Flags().BoolVar(&options.GlobalFlagHumanReadable, "human", false, "output human readable info")

}
