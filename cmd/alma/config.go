package main

type Config struct {
	URL    string `mapstructure:"url"`
	ApiKey string `mapstructure:"api_key"`
}
