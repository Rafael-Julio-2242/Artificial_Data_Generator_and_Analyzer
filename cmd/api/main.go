package main

import (
	"artificial-data-analyzer-generation/internal/adapters/fileops"
	httplayer "artificial-data-analyzer-generation/internal/adapters/http_layer"
	"artificial-data-analyzer-generation/internal/adapters/http_layer/middlewares"
	"artificial-data-analyzer-generation/internal/adapters/services"
	"context"
	"log"

	"github.com/joho/godotenv"
	"google.golang.org/genai"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error on loading enviroment: %v", err)
	}

	ctx := context.Background()

	aiClient, err := genai.NewClient(ctx, nil)

	if err != nil {
		log.Fatalf("Error on loading AI client: %v", err)
	}

	fileParser := fileops.NewFileParser()
	csvGenerator := fileops.NewCSVGenerator()
	defineVariableTypesService := services.NewDefineVariableTypesService(aiClient, ctx)
	dataCleaningService := services.NewDataCleaningService()
	calculateCentralTrendsService := services.NewCalculateCentralTrendsService()
	calculateFrequenciesService := services.NewCalculateFrequenciesService()

	server := httplayer.GetHttpServer(
		middlewares.InjectFileParserMiddleware(fileParser),
		middlewares.InjectCSVGeneratorMiddleware(csvGenerator),
		middlewares.InjectDefineVariableTypesMiddleware(defineVariableTypesService),
		middlewares.InjectDataCleaningMiddleware(dataCleaningService),
		middlewares.InjectCalculateCentralTrendsMiddleware(calculateCentralTrendsService),
		middlewares.InjectCalculateFrequenciesMiddleware(calculateFrequenciesService),
	)

	server.Run(":8080")
}
