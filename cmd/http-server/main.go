package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Rikypurnomo/warmup/config"
	"github.com/Rikypurnomo/warmup/pkg/cache"
	"github.com/Rikypurnomo/warmup/pkg/database"
	postgres "github.com/Rikypurnomo/warmup/pkg/database"
	"github.com/Rikypurnomo/warmup/pkg/logger"
	"github.com/Rikypurnomo/warmup/server"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "warmup"
	app.Usage = "run warmup-server"
	app.Version = fmt.Sprintf("%s built on %s (commit: %s)", Version, BuildDate, Commit)
	app.Description = "warmup is a web server for microservice architecture"
	app.Commands = []cli.Command{
		{
			Name:        "start",
			Description: "Start warmup server",
			Action:      startKiyora,
		},
		{
			Name:        "migrate",
			Description: "Migrate database",
			Action:      migrateData,
		},
	}

	app.Run(os.Args)
}

func migrateData(_ *cli.Context) error {
	config.Load()
	logger.SetupLogger()

	database.Connect()
	defer database.CloseConnect()

	database.Migrate(database.MigrateList...)
	return nil
}

func startKiyora(_ *cli.Context) error {
	sigC := make(chan os.Signal, 1)
	signal.Notify(sigC, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	config.Load()
	logger.SetupLogger()

	ginLoad := server.NewRouter()
	ginLoad.RoutersConfig()

	// uncomment if you are have module adapter/az
	// if config.IsEnabledAdapterAz() {
	// 	az.LoadEnviroment(config.AdapterAzHost())
	// 	defer az.CloseConnection()
	// }

	if config.IsEnabledPG() {
		postgres.Connect()
		defer postgres.CloseConnect()
	}

	if config.IsEnabledRedis() {
		cache.InitCache()
		defer cache.CloseConnectionCache()
	}

	ctx, cancel := context.WithCancel(context.Background())
	go server.StartServer(ctx, ginLoad.Router)

	sig := <-sigC
	logger.Debugf("Received %d, shutting down", sig)

	defer cancel()
	server.ShutdownServer(ctx)

	return nil
}

var (
	Version   = "1.0.0"
	Commit    = "1.0"
	BuildDate = "2023-06-08"
)
