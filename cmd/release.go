package cmd

import (
	"fmt"
	"github.com/GoFarsi/gvm/api"
	"github.com/GoFarsi/gvm/errors"
	"log"

	"github.com/spf13/cobra"
)

// releaseCmd represents the release command
var releaseCmd = &cobra.Command{
	Use:   "release",
	Short: "Show version change logs",
	Long: `Show version change logs of go versions.

For example:
  $ gvm release
  $ gvm release --version 1.x`,
	Run: func(cmd *cobra.Command, args []string) {
		ver, err := cmd.Flags().GetString("version")
		if err != nil {
			log.Fatalln(errors.ERR_INVALID_VALUE)
		}

		if len(ver) != 0 {
			fmt.Println(api.GetDocVersion(ver))
			return
		}

		fmt.Println(api.GetFullDoc())
	},
}

func init() {
	rootCmd.AddCommand(releaseCmd)

	releaseCmd.Flags().StringP("version", "v", "", "set specific version for download")
}
