package main

type Config struct {
	Alma struct {
		URL    string `env:"URL,required"`
		ApiKey string `env:"API_KEY,required"`
	} `envPrefix:"ALMA_"`
}
