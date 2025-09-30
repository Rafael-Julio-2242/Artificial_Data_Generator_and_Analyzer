package handlers

import (
	"artificial-data-analyzer-generation/internal/domain/ports"

	"github.com/gin-gonic/gin"
)

func DefineVariableTypes(context *gin.Context) {

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

	if err != nil {
		context.JSON(500, gin.H{"message": "error getting file parser"})
		return
	}

	// fname := fileHeader.Filename
	mtype := fileHeader.Header.Get("Content-Type")

	defer file.Close()

	switch mtype {
	case "text/csv":
		data, err := fileParser.ConvertFileToData(&file, "csv")
		if err != nil {
			context.JSON(500, gin.H{"message": "error converting file to data"})
			return
		}
		context.JSON(200, gin.H{"data": data})
	case "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet":
		data, err := fileParser.ConvertFileToData(&file, "xlsx")
		if err != nil {
			context.JSON(500, gin.H{"message": "error converting file to data"})
			return
		}
		context.JSON(200, gin.H{"data": data})
	default:
		context.JSON(400, gin.H{"message": "file type not supported"})
		return
	}

}
