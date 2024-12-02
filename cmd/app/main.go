package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/apriSinggih/movie-app/config"
	"github.com/apriSinggih/movie-app/internal/builder"
	"github.com/apriSinggih/movie-app/pkg/database"
	server "github.com/apriSinggih/movie-app/pkg/sever"
)

func main() {
	cfg, err := config.NewConfig(".env")
	checkError(err)
	db, err := database.InitDatabase(cfg.MySQLConfig)
	checkError(err)

	publicRoutes := builder.BuildPublicRoutes(cfg, db)
	privateRoutes := builder.BuildPrivateRoutes(db)

	srv := server.NewServer(publicRoutes, privateRoutes)
	runServer(srv, cfg.PORT)
	waitForShutdown(srv)
}

func waitForShutdown(srv *server.Server) {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	go func() {
		if err := srv.Shutdown(ctx); err != nil {
			srv.Logger.Fatal(err)
		}
	}()
}

func runServer(srv *server.Server, port string) {
	go func() {
		err := srv.Start(fmt.Sprintf(":%s", port))
		log.Fatal(err)
	}()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
