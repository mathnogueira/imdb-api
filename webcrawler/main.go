package main

import (
	"strings"

	"github.com/mathnogueira/imdb-api/webcrawler/crawler"
	"github.com/spf13/viper"
)

func main() {
	config, err := getConfig()

	if err != nil {
		panic(err)
	}

	crawlerOptions := crawler.Options{StorageURL: config.GetString("storage.url")}
	err = crawler.Execute(crawlerOptions)

	if err != nil {
		panic(err)
	}
}

func getConfig() (*viper.Viper, error) {
	config := viper.New()

	config.SetEnvPrefix("crawler")
	config.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	config.AutomaticEnv()

	return config, nil
}
