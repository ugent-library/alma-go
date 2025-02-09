package main

import (
	"github.com/spf13/cobra"
	"github.com/tidwall/pretty"
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
