package main

import (
	"context"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var config Config

func main() {
	v := viper.New()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.BindEnv("alma.url")
	v.BindEnv("alma.api_key")
	cobra.CheckErr(v.Unmarshal(&config))

	cobra.CheckErr(rootCmd.ExecuteContext(context.Background()))
}
