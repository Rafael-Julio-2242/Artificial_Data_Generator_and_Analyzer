package main

import (
	"artificial-data-analyzer-generation/internal/adapters/fileops"
	httplayer "artificial-data-analyzer-generation/internal/adapters/http_layer"
	"artificial-data-analyzer-generation/internal/adapters/http_layer/middlewares"
)

func main() {

	fileParser := fileops.NewFileParser()
	server := httplayer.GetHttpServer(
		middlewares.InjectFileParserMiddleware(fileParser),
	)

	server.Run(":8080")
}
