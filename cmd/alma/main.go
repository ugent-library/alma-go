package main

import (
	"context"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ugent-library/alma-go"
)

var config Config

var almaClient *alma.Client

func main() {
	v := viper.New()
	v.SetEnvKeyReplacer(strings.NewReplacer(`.`, `_`))
	v.SetEnvPrefix("alma")
	v.BindEnv("url")
	v.BindEnv("api_key")
	cobra.CheckErr(v.Unmarshal(&config))

	var err error
	almaClient, err = alma.New(alma.Config{
		URL:    config.URL,
		ApiKey: config.ApiKey,
	})
	cobra.CheckErr(err)

	cobra.CheckErr(rootCmd.ExecuteContext(context.Background()))
}
