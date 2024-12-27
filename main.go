package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"seeyou-go/api/router"
	"seeyou-go/config"
	"syscall"
	"time"
)

func main() {

	config.InitConfig()

	port := config.AppConfig.App.Port

	r := router.SetupRouter()

	// 提供 uploads 目录的静态文件服务
	r.Static("/uploads", "./uploads")

	srv := &http.Server{
		Addr:    port,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")

}
