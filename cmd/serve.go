package cmd

import (
	"github.com/randijulio13/gofiber/internal/config"
	"github.com/randijulio13/gofiber/internal/router"
	"github.com/randijulio13/gofiber/pkg/app"
	"github.com/randijulio13/gofiber/pkg/database"
	"github.com/randijulio13/gofiber/pkg/logger"
	customValidator "github.com/randijulio13/gofiber/pkg/validator"

	"github.com/spf13/cobra"
)

func startServer(cmd *cobra.Command, args []string) {
	app.SetConfig(config.GetFiberConfig())

	logger.InitializeLogger()
	customValidator.InitializeValidator()

	config := config.GetConfig()
	dsn := database.GetDsn(config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)
	database.OpenConnection(dsn)

	if appServer := app.InitializeServer(); appServer != nil {
		router.Route(appServer)
	}

	app.StartServer(config.AppPort)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start fiber appserver.",
	Run:   startServer,
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
