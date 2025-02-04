package main

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ugent-library/alma-go"
)

var config Config

var almaClient *alma.Client

func main() {
	viper.SetEnvPrefix("alma")
	viper.BindEnv("url")
	viper.BindEnv("api_key")
	cobra.CheckErr(viper.Unmarshal(&config))

	var err error
	almaClient, err = alma.New(alma.Config{
		URL:    config.URL,
		ApiKey: config.ApiKey,
	})
	cobra.CheckErr(err)

	cobra.CheckErr(rootCmd.ExecuteContext(context.Background()))
}
