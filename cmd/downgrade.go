package cmd

import (
	"context"
	"github.com/GoFarsi/gvm/internal"
	"github.com/spf13/cobra"
	"log"
)

// downgradeCmd represents the downgrade command
var downgradeCmd = &cobra.Command{
	Use:   "downgrade",
	Short: "Downgrade go version to previous version",
	Long: `With this command you can downgrade previous version of golang
or specific version x.x.x.

For example:
  $ gvm downgrade
  $ gvm downgrade --backup /home/user
  $ gvm downgrade --version 1.x.x
  $ gvm downgrade --version 1.x.x --backup /home/user
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

		if err := internal.DowngradeGo(context.Background(), ver, backup); err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(downgradeCmd)

	downgradeCmd.Flags().StringP("backup", "b", "", "backup old version installed with downloaded golang into specific path")
	downgradeCmd.Flags().StringP("version", "v", "", "set specific version for download")
}
