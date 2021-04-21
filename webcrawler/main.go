package main

import (
	"strings"

	"github.com/mathnogueira/imdb-api/webcrawler/extractor"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	config, err := getConfig()

	if err != nil {
		panic(err)
	}

	logger, err := zap.NewDevelopment(zap.Fields(
		zap.String("app", "crawler"),
	))
	if err != nil {
		panic(err)
	}

	crawlerOptions := extractor.Options{StorageURL: config.GetString("storage.url")}
	extractor := extractor.NewExtractor(logger)
	err = extractor.Execute(crawlerOptions)

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
