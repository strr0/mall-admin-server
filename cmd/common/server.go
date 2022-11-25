package common

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"log"
	"mall-admin-server/common/router"
	"mall-admin-server/config"
	"mall-admin-server/docs"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	redis, mongo, rabbit bool
	StartCmd             = &cobra.Command{
		Use:   "common",
		Short: "common",
		Long:  "common",
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func init() {
	StartCmd.Flags().BoolVar(&redis, "redis", false, "--redis=true|false")
	StartCmd.Flags().BoolVar(&mongo, "mongo", false, "--mongo=true|false")
	StartCmd.Flags().BoolVar(&rabbit, "rabbit", false, "--rabbit=true|false")
}

func initConfig() ([]func(), []func()) {
	start := make([]func(), 0)
	end := make([]func(), 0)
	if redis {
		start = append(start, config.InitRedis)
		end = append(end, config.CloseRedis)
	}
	if mongo {
		start = append(start, config.InitMongo)
		end = append(end, config.CLoseMongo)
	}
	if rabbit {
		start = append(start, config.InitRabbit)
		end = append(end, config.CloseRabbit)
	}
	return start, end
}

// 路由
func initRouter(e *gin.Engine) {
	router.InitRouter(e)
	docs.InitSwagger(e) // 引入swagger
}

func run() {
	start, end := initConfig()
	for _, f := range start {
		f()
	}
	for _, df := range end {
		defer df()
	}
	e := gin.Default()
	//initRouter(e)
	srv := http.Server{
		Addr:    ":8080",
		Handler: e,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	// 等待终止信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}
	log.Println("Server exiting")
}
