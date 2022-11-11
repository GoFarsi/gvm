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
  $ gvm upgrade --backup
  $ gvm upgrade --version 1.x.x
  $ gvm upgrade --version 1.x.x --backup
`,
	Run: func(cmd *cobra.Command, args []string) {
		ver, err := cmd.Flags().GetString("version")
		if err != nil {
			log.Fatalln(err)
		}

		backup, err := cmd.Flags().GetBool("backup")
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

	upgradeCmd.Flags().BoolP("backup", "b", false, "backup old version installed with downloaded golang into /home/{user}")
	upgradeCmd.Flags().StringP("version", "v", "", "set specific version for download")
}
