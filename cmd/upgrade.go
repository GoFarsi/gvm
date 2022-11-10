package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// upgradeCmd represents the upgrade command
var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade go to last version of released",
	Long: `With this command you can upgrade installed golang to last version 
or specific version x.x.x.

For example:
  $ gvm upgrade
  $ gvm upgrade --backup
  $ gvm upgrade --version 1.x.x
  $ gvm upgrade --version 1.x.x --backup
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("upgrade called")
	},
}

func init() {
	rootCmd.AddCommand(upgradeCmd)

	upgradeCmd.Flags().BoolP("backup", "b", false, "backup old version installed with downloaded golang into /home/{user}")
	upgradeCmd.Flags().StringP("version", "v", "", "set specific version for download")
}
