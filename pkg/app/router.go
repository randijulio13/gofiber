package app

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

var appServer *fiber.App
var appServerConfig fiber.Config

func SetConfig(config fiber.Config) {
	appServerConfig = config
}

func InitializeServer() *fiber.App {
	appServer = fiber.New(appServerConfig)
	return appServer
}

func StartServer(port string) {
	appServer.Listen(fmt.Sprintf(":%s", port))
}

func GetAppServer() *fiber.App {
	return appServer
}
