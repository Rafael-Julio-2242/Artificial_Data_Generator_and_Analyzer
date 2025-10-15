package handlers

import (
	"artificial-data-analyzer-generation/internal/domain/ports"
	"artificial-data-analyzer-generation/internal/domain/services"
	"strings"

	"github.com/gin-gonic/gin"
)

func DataCleaning(context *gin.Context) {

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

	switch mtype {
	case "text/csv":
		data, err := fileParser.ConvertFileToData(&file, "csv")
		if err != nil {
			context.JSON(500, gin.H{"message": "error converting file to data"})
			return
		}

		isCsv := strings.HasSuffix(fname, ".csv")

		if !isCsv {
			context.JSON(400, gin.H{"message": "mimetype does not correspond to file suffix!"})
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

		context.JSON(200, gin.H{"data": cleanedData})
		return
	case "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet":
		data, err := fileParser.ConvertFileToData(&file, "xlsx")
		if err != nil {
			context.JSON(500, gin.H{"message": "error converting file to data"})
			return
		}

		isXlsx := strings.HasSuffix(fname, ".xlsx")

		if !isXlsx {
			context.JSON(400, gin.H{"message": "mimetype does not correspond to file suffix!"})
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

		context.JSON(200, gin.H{"data": cleanedData})
		return
	default:
		context.JSON(400, gin.H{"message": "file type not supported"})
		return

	}

}
