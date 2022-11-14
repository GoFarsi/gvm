package cmd

import (
	"github.com/GoFarsi/gvm/internal"
	"github.com/spf13/cobra"
	"log"
)

// notificationCmd represents the notification command
var notificationCmd = &cobra.Command{
	Use:   "notification",
	Short: "check new version of golang",
	Long: `check new version of golang base on installed version, for automatic check you can set cronjob on you OS.
	
	Example cron for every 6 hours check new version:
	(crontab -l ; echo "* */6 * * * gvm notification >/dev/null 2>&1")| crontab -
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := internal.ReleaseNotification(); err != nil {
			log.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(notificationCmd)
}
