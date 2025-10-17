package handlers

import (
	"artificial-data-analyzer-generation/internal/domain/ports"
	"artificial-data-analyzer-generation/internal/domain/services"

	"github.com/gin-gonic/gin"
)

func CalculateCentralTrends(context *gin.Context) {
	fileHeader, err := context.FormFile("file")

	if err != nil {
		context.JSON(400, gin.H{"message": "file is required"})
		return
	}

	file, err := fileHeader.Open()

	if err != nil {
		context.JSON(500, gin.H{"message": "error opening file"})
		return
	}

	fileParserSv, ok := context.Get(ports.FileParserServiceKey)

	if !ok {
		context.JSON(500, gin.H{"message": "error getting file parser"})
		return
	}

	fileParser := fileParserSv.(ports.FileParser)

	fname := fileHeader.Filename
	mtype := fileHeader.Header.Get("Content-Type")

	defer file.Close()

	data, err := fileParser.ConvertFileToData(&file, mtype, fname)

	if err != nil {
		context.JSON(500, gin.H{"message": "error converting file to data: " + err.Error()})
		return
	}

	dataCleaningSv, ok := context.Get(services.DataCleaningServiceKey)

	if !ok {
		context.JSON(500, gin.H{"message": "error getting data cleaning service"})
		return
	}

	dataCleaningService := dataCleaningSv.(services.DataCleaningService)

	cleanedData, err := dataCleaningService.CleanData(data)

	if err != nil {
		context.JSON(500, gin.H{"message": "error cleaning data: " + err.Error()})
		return
	}

	defineVariableTypesServiceSv, ok := context.Get(services.DefineVariableTypesServiceKey)

	if !ok {
		context.JSON(500, gin.H{"message": "error getting define variable types service"})
		return
	}

	defineVariableTypesService := defineVariableTypesServiceSv.(services.DefineVariableTypesService)

	fixedData, err := defineVariableTypesService.DefineVariableTypes(cleanedData)

	if err != nil {
		context.JSON(500, gin.H{"message": "error defining variable types: " + err.Error()})
		return
	}

	calculateCentralTrendsServiceSv, ok := context.Get(services.CalculateCentralTrendsServiceKey)

	if !ok {
		context.JSON(500, gin.H{"message": "error getting calculate central trends service"})
		return
	}

	calculateCentralTrendsService := calculateCentralTrendsServiceSv.(services.CalculateCentralTrends)

	centralTrendsData, err := calculateCentralTrendsService.Calculate(fixedData)

	if err != nil {
		context.JSON(500, gin.H{"message": "error calculating central trends: " + err.Error()})
		return
	}

	context.JSON(200, gin.H{"data": centralTrendsData})
}
