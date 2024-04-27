/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"energy-response-assignment/config"
	"energy-response-assignment/util"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
)

// migrateUpCmd represents the migrateUp command
var migrateUpCmd = &cobra.Command{
	Use:   "up",
	Short: "migrate to v1 command",
	Long:  `Command to install version 1 of our application`,
	Run: func(cmd *cobra.Command, args []string) {
		conf, err := util.GetConfig[config.Config]()
		if err != nil {
			panic(err)
		}
		m, err := migrate.New("file://db/migrations", conf.Database.GetConnectionString())
		if err != nil {
			panic(err)
		}
		if err := m.Up(); err != nil {
			panic(err)
		}

		version, _, err := m.Version()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Finish Migrate Up to version: %d", version)
	},
}

func init() {
	migrateCmd.AddCommand(migrateUpCmd)
}
