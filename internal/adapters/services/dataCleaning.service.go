package services

import (
	"strings"
)

type DataCleaningService struct{}

func NewDataCleaningService() *DataCleaningService {
	return &DataCleaningService{}
}

func (d *DataCleaningService) CleanData(data map[string][]string) (map[string][]string, error) {
	// Aqui eu preciso limpar dados incompletos e dados repetidos
	var resultData map[string][]string = make(map[string][]string)
	var dataMap map[string]bool = make(map[string]bool)
	var headers []string = make([]string, 0)

	var qtd int = 0

	// Primeira Limpeza: Remover dados duplicados
	// Primeiro passo - Pegar todos os headers
	for header, values := range data {
		headers = append(headers, header)

		if len(values) > qtd {
			qtd = len(values)
		}
	}

	// Segundo passo - Popular "dataMap" com indexação de informações
	for index := 0; index < qtd; index++ {
		var dataMapValue string = ""

		for _, header := range headers {
			dataMapValue += data[header][index] + "-"
		}

		if _, ok := dataMap[dataMapValue]; ok {
			continue
		}

		dataMap[dataMapValue] = true
	}

	// Terceiro passo - Popular o "result data" com os dados não duplicados armazenados como chave no dataMap
	for valueKey := range dataMap {

		valueKey = strings.TrimSuffix(valueKey, "-")

		values := strings.Split(valueKey, "-")

		for index, header := range headers {
			resultData[header] = append(resultData[header], values[index])
		}

	}

	// Segunda Limpeza: Remover dados incompletos
	for _, values := range resultData {
		for index, value := range values {
			if value == "" {
				// Eu preciso apagar isso de todos os outros headers também
				for _, header := range headers {
					resultData[header] = append(resultData[header][:index], resultData[header][index+1:]...)
				}
			}
		}
	}

	return resultData, nil
}
