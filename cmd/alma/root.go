package main

import (
	"github.com/spf13/cobra"
	"github.com/tidwall/pretty"
	"github.com/ugent-library/alma-go"
)

var prettify = false

var rootCmd = &cobra.Command{
	Use:   "alma",
	Short: "",
	Long:  ``,
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&prettify, "pretty", false, "")
}

func newAlmaClient() *alma.Client {
	client, err := alma.New(alma.Config{
		URL:    config.URL,
		ApiKey: config.ApiKey,
	})
	cobra.CheckErr(err)
	return client
}

func writeString(cmd *cobra.Command, str string) error {
	_, err := cmd.OutOrStdout().Write([]byte(str))
	return err
}

func writeJSON(cmd *cobra.Command, b []byte) error {
	if prettify {
		b = pretty.Color(pretty.Pretty(b), pretty.TerminalStyle)
	}
	_, err := cmd.OutOrStdout().Write(b)
	return err
}
