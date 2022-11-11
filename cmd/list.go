package cmd

import (
	"fmt"
	"github.com/GoFarsi/gvm/errors"
	"github.com/GoFarsi/gvm/internal"
	"log"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List of go version with commit id",
	Long: `With this command you can get list of golang version.

For example:
  $ gvm list
  $ gvm list --line 10`,
	Run: func(cmd *cobra.Command, args []string) {
		list, err := internal.NewList()
		if err != nil {
			log.Fatalln(errors.ERR_NETWORK_TIMEOUT)
		}

		line, err := cmd.Flags().GetInt("line")
		if err != nil {
			log.Fatalln(errors.ERR_INVALID_VALUE)
		}

		if line != 0 {
			list.SetNumOfVersions(line)
		}

		fmt.Println(list.Print())
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().IntP("line", "l", 0, "limit to n of last version list")
}
