package main

import "github.com/mathnogueira/imdb-api/storage/api"

func main() {
	server := api.NewServer(8000)
	server.Start()
}
