/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/randijulio13/gofiber/internal/config"
	"github.com/randijulio13/gofiber/pkg/logger"
	"github.com/randijulio13/gofiber/pkg/migration"

	"github.com/pressly/goose/v3"
	"github.com/spf13/cobra"
)

var (
	dir     string
	version bool
	verbose bool
	table   string
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		config.InitializeConfig()
		logger.InitializeLogger()

		if version {
			migration.ShowVersion()
			return
		}

		goose.SetVerbose(verbose)

		if len(args) > 0 {
			migration.Run(cmd, args)
		}
	},
}

func init() {
	flags := migrateCmd.Flags()
	flags.StringVarP(&dir, "dir", "d", "migration", "set database migration sql folder path")
	flags.StringVarP(&table, "table", "t", "migration", "set goose table migration name")
	flags.BoolVar(&version, "version", false, "print goose version")
	flags.BoolVar(&verbose, "verbose", false, "set goose verbose")
	rootCmd.AddCommand(migrateCmd)
}
