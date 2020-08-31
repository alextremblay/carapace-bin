package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var unpauseCmd = &cobra.Command{
	Use:   "unpause",
	Short: "Unpause services",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.AddCommand(unpauseCmd)

	carapace.Gen(unpauseCmd).PositionalAnyCompletion(ActionServices())
}
