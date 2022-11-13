package cmd

import (
	"context"
	"github.com/GoFarsi/gvm/internal"
	"log"

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
  $ gvm upgrade --backup /home/user
  $ gvm upgrade --version 1.x.x
  $ gvm upgrade --version 1.x.x --backup /home/user
`,
	Run: func(cmd *cobra.Command, args []string) {
		ver, err := cmd.Flags().GetString("version")
		if err != nil {
			log.Fatalln(err)
		}

		backup, err := cmd.Flags().GetString("backup")
		if err != nil {
			log.Fatalln(err)
		}

		if err := internal.UpgradeGo(context.Background(), ver, backup); err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(upgradeCmd)

	upgradeCmd.Flags().StringP("backup", "b", "", "backup old version installed with downloaded golang into specific path")
	upgradeCmd.Flags().StringP("version", "v", "", "set specific version for download")
}
