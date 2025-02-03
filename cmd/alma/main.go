package main

import (
	"log"
	"os"

	"github.com/caarlos0/env/v11"
	"github.com/ugent-library/alma-go"
)

var almaClient *alma.Client

func main() {
	var config Config

	if err := env.Parse(&config); err != nil {
		log.Fatal(err)
	}

	almaClient = alma.New(alma.Config{
		URL:    config.Alma.URL,
		ApiKey: config.Alma.ApiKey,
	})

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
