package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"mall-admin-server/cmd/admin"
	"os"
)

var rootCmd = &cobra.Command{
	Use:          "mall-admin-server",
	Short:        "mall-admin-server",
	Long:         "mall-admin-server",
	SilenceUsage: true,
	Run: func(cmd *cobra.Command, args []string) {
		tip()
	},
}

func init() {
	rootCmd.AddCommand(admin.StartCmd)
}

func tip() {
	log.Println("mall-admin-server running...")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println("running failed.")
		os.Exit(-1)
	}
}
