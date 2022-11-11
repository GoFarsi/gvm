package cmd

import (
	"context"
	"github.com/GoFarsi/gvm/internal"
	"log"

	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install last version of golang",
	Long: `With this command you can install last version of golang
or specific version x.x.x.

For example:
  $ gvm install
  $ gvm install --backup
  $ gvm install --version 1.x.x
  $ gvm install --version 1.x.x --backup
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

		if err := internal.InstallGo(context.Background(), ver, backup); err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	installCmd.Flags().BoolP("backup", "b", false, "backup downloaded golang into /home/{user}")
	installCmd.Flags().StringP("version", "v", "", "set specific version for download")
}
