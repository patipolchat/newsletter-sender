/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"energy-response-assignment/config"
	"energy-response-assignment/util"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/spf13/cobra"
)

// migrateDownCmd represents the migrateDown command
var migrateDownCmd = &cobra.Command{
	Use:   "down",
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
		if err := m.Down(); err != nil {
			panic(err)
		}
		version, _, err := m.Version()
		if err != nil {
			fmt.Println("Error getting version after migration down")
			return
		}
		fmt.Printf("Finish Migrate Down to version: %d", version)
	},
}

func init() {
	migrateCmd.AddCommand(migrateDownCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// migrateDownCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// migrateDownCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
