package generator

import (
	"fmt"
	"github.com/spf13/cobra"
	"gorm.io/gen"
	"log"
	"mall-admin-server/config"
)

var StartCmd = &cobra.Command{
	Use:   "generator",
	Short: "generator",
	Long:  "generator",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			log.Println("Usage: generator [pkg] [table]")
			return
		}
		generate(args[0], args[1])
	},
}

func generate(pkg, table string) {
	g := gen.NewGenerator(gen.Config{
		OutPath: fmt.Sprintf("./%s/query", pkg),
		Mode:    gen.WithoutContext | gen.WithDefaultQuery,
	})
	db := config.GetDb()
	g.UseDB(db)
	g.GenerateModel(table)
	g.Execute()
}
