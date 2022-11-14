package cmd

import (
	"fmt"
	"github.com/GoFarsi/gvm/internal"
	"github.com/GoFarsi/gvm/internal/gvm"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "gvm",
	Version: internal.GvmVersion.String(),
	Short:   "go version manager tool",
	Long: `You can manage your golang versions easily 
with gvm by installing, upgrading, list and other commands.
`,
	Run: func(cmd *cobra.Command, args []string) {
		checkUpdate, err := cmd.Flags().GetBool("update")
		if err != nil {
			log.Fatalln(err)
		}

		if checkUpdate {
			gvm, err := gvm.NewGVMCheck()
			if err != nil {
				log.Fatalln(err)
			}

			msg, err := gvm.CheckNewVersion()
			if err != nil {
				log.Fatalln(err)
			}

			fmt.Println(msg)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("update", "u", false, "check new version of gvm")
}
