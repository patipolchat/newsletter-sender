/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"energy-response-assignment/app/repository"
	"energy-response-assignment/app/service"
	"energy-response-assignment/config"
	"energy-response-assignment/util"
	"energy-response-assignment/util/db"
	"energy-response-assignment/util/mailer"
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

var (
	header string
	body   string
)

// sendNewsletterCmd represents the sendNewsletter command
var sendNewsletterCmd = &cobra.Command{
	Use:   "sendNewsletter",
	Short: "sendNewsletter to active subscriber",
	Run: func(cmd *cobra.Command, args []string) {
		if header == "" || body == "" {
			fmt.Println("Header and body are required")
			return
		}
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
		mailDialer := mailer.NewMailer(&mailer.Config{
			Host:     cfg.Mailer.Host,
			Port:     cfg.Mailer.Port,
			Username: cfg.Mailer.Username,
			Password: cfg.Mailer.Password,
			From:     cfg.Mailer.From,
		})
		repo := repository.NewNewsletter(cfg, pool.GetPool())
		newsletter := service.NewNewsletter(cfg, repo, mailDialer)
		ctx, cancel := context.WithTimeout(cmd.Context(), 100*time.Second)
		defer cancel()
		err = newsletter.SendNewsLetter(ctx, header, body)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(sendNewsletterCmd)
	sendNewsletterCmd.PersistentFlags().StringVar(&header, "header", "", "Header of the newsletter")
	sendNewsletterCmd.PersistentFlags().StringVar(&body, "body", "", "Body of the newsletter")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sendNewsletterCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sendNewsletterCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
