/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"energy-response-assignment/app"
	"energy-response-assignment/config"
	"energy-response-assignment/server"
	"energy-response-assignment/util"
	"energy-response-assignment/util/db"
	"github.com/spf13/cobra"
)

// serveApiCmd represents the serveApi command
var serveApiCmd = &cobra.Command{
	Use:   "serveApi",
	Short: "Serve the API",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := util.GetConfig[config.Config]()
		if err != nil {
			panic(err)
		}
		pool, err := db.NewPgxPool(cfg.Database.GetConnectionString())
		if err != nil {
			panic(err)
		}
		defer pool.Disconnect()
		if pool.Ping() != nil {
			panic("Database connection failed")
		}
		server := server.NewServer(cfg)
		app := app.NewApp(cfg, pool.GetPool(), server)
		app.Start()
	},
}

func init() {
	rootCmd.AddCommand(serveApiCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveApiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveApiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
