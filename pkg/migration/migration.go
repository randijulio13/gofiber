package migration

import (
	"fmt"

	"github.com/randijulio13/gofiber/internal/config"
	"github.com/randijulio13/gofiber/pkg/database"
	"github.com/randijulio13/gofiber/pkg/logger"

	"github.com/pressly/goose/v3"
	"github.com/spf13/cobra"
)

func ShowVersion() {
	fmt.Println(goose.VERSION)
}

func Run(cmd *cobra.Command, args []string) {
	log := logger.GetLogger()
	config := config.GetConfig()

	dir := cmd.Flag("dir").Value.String()
	table := cmd.Flag("table").Value.String()
	goose.SetTableName(table)

	command := args[0]
	switch command {
	case "create":
		if err := goose.Run(command, nil, dir, args[1:]...); err != nil {
			log.Fatalf("goose run: %v", err)
		}
		return
	case "fix":
		if err := goose.Run(command, nil, dir, args[1:]...); err != nil {
			log.Fatalf("goose run: %v", err)
		}
		return
	}

	db, err := goose.OpenDBWithDriver("mysql", database.GetDsn(config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName))
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("db migrate: failed to close DB: %v\n", err)
		}
	}()

	arguments := []string{}
	if len(args) > 3 {
		arguments = append(arguments, args[3:]...)
	}

	if err := goose.Run(command, db, dir, arguments...); err != nil {
		log.Fatalf("db migrate run: %v", err)
	}
}
