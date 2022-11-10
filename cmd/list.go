package cmd

import (
	"fmt"
	"github.com/GoFarsi/gvm/api"
	"github.com/GoFarsi/gvm/errors"
	"log"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list of go version with commit id",
	Run: func(cmd *cobra.Command, args []string) {
		list, err := api.NewList()
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
	listCmd.Flags().IntP("line", "l", 0, "limit to n of last version list")
}
