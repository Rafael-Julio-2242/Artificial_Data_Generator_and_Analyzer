package handlers

import (
	"artificial-data-analyzer-generation/internal/domain/ports"
	"artificial-data-analyzer-generation/internal/domain/services"
	"strings"

	"github.com/gin-gonic/gin"
)

// Aqui eu preciso definir os tipos de dados das variáveis / colunas
// Preciso definir se são:
//	Qualitativa Nominal - Texto normal sem hierarquia
//	Qualitativa Ordinal - Texto com algum tipo de Hierarquia (menor valor, maior valor)
//	Quantitativa Discreta - Números inteiros
//	Quantitativa Contínua - Números com pontos flutuantes

func DefineVariableTypes(context *gin.Context) { // TODO Ainda preciso ajustar isso

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

		// Depois posso fazer uma verificação de bytes, por enquanto isso é suficiente

		// Aqui eu preciso fazer a verificação de tipos, definir quais os tipos de variáveis

		defineVariableTypesServiceSv, ok := context.Get(services.DefineVariableTypesServiceKey)

		if !ok {
			context.JSON(500, gin.H{"message": "error getting define variable types service"})
			return
		}

		defineVariableTypesService := defineVariableTypesServiceSv.(services.DefineVariableTypesService)

		fixedData, err := defineVariableTypesService.DefineVariableTypes(data)

		if err != nil {
			context.JSON(500, gin.H{"message": "error defining variable types: " + err.Error()})
			return
		}

		context.JSON(200, gin.H{"data": fixedData})
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

		// Depois posso fazer uma verificação de bytes

		defineVariableTypesServiceSv, ok := context.Get(services.DefineVariableTypesServiceKey)

		if !ok {
			context.JSON(500, gin.H{"message": "error getting define variable types service"})
			return
		}

		defineVariableTypesService := defineVariableTypesServiceSv.(services.DefineVariableTypesService)

		fixedData, err := defineVariableTypesService.DefineVariableTypes(data)

		if err != nil {
			context.JSON(500, gin.H{"message": "error defining variable types: " + err.Error()})
			return
		}

		context.JSON(200, gin.H{"data": fixedData})
		return
	default:
		context.JSON(400, gin.H{"message": "file type not supported"})
		return
	}

}
