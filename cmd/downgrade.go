package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// downgradeCmd represents the downgrade command
var downgradeCmd = &cobra.Command{
	Use:   "downgrade",
	Short: "Downgrade go version to previous version",
	Long: `With this command you can downgrade previous version of golang
or specific version x.x.x.

For example:
  $ gvm downgrade
  $ gvm downgrade --backup
  $ gvm downgrade --version 1.x.x
  $ gvm downgrade --version 1.x.x --backup
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("implement me!!")
	},
}

func init() {
	rootCmd.AddCommand(downgradeCmd)

	downgradeCmd.Flags().BoolP("backup", "b", false, "backup old version installed with downloaded golang into /home/{user}")
	downgradeCmd.Flags().StringP("version", "v", "", "set specific version for download")
}
