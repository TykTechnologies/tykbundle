package cmd

import (
	"github.com/TykTechnologies/tyk-cli/cmd/usage"
	"github.com/TykTechnologies/tyk-cli/commands/remote"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add remotes to configuration",
	Long:  `Use this command to add remotes to configuration and post new organisations to Tyk`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			remote.Add("example.conf.json", args)
			return
		}
		cmd.Usage()
	},
}

func init() {
	remoteCmd.AddCommand(addCmd)
	usage.Add(addCmd)
}
