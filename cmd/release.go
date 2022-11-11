package cmd

import (
	"fmt"
	"github.com/GoFarsi/gvm/errors"
	"github.com/GoFarsi/gvm/internal"
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
			fmt.Println(internal.GetDocWithSpecificVersion(ver))
			return
		}

		doc, err := internal.GetFullDoc()
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(doc)
	},
}

func init() {
	rootCmd.AddCommand(releaseCmd)

	releaseCmd.Flags().StringP("version", "v", "", "set specific version for download")
}
