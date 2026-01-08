package main

type Config struct {
	Alma struct {
		URL    string `mapstructure:"url"`
		ApiKey string `mapstructure:"api_key"`
	} `mapstructure:"alma"`
}
