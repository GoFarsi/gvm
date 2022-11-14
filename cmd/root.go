package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var Version = "1.1.0"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "gvm",
	Version: Version,
	Short:   "go version manager tool",
	Long: `You can manage your golang versions easily 
with gvm by installing, upgrading, list and other commands.
`,
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
}
