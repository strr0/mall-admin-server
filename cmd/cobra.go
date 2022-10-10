package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"log"
	"mall-admin-server/cmd/admin"
	"mall-admin-server/cmd/generator"
	"os"
)

var rootCmd = &cobra.Command{
	Use:          "mall-admin-server",
	Short:        "mall-admin-server",
	Long:         "mall-admin-server",
	SilenceUsage: true,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			log.Println("Usage: [cmd]")
			return errors.New("requires at least one arg")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		tip()
	},
}

func init() {
	rootCmd.AddCommand(admin.StartCmd)
	rootCmd.AddCommand(generator.StartCmd)
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
