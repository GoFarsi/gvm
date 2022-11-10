package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove golang completely with cache files",
	Long: `Remove golang completely with cache files.

For example:
  $ gvm remove
  $ gvm remove --all`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("implement me!!")
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)

	installCmd.Flags().BoolP("all", "a", false, "remove golang with other files (etc cache and ...)")
}
