package services

import (
	"errors"
	"fmt"
	"strings"
)

type CalculateCentralTrendsService struct{}

func NewCalculateCentralTrendsService() *CalculateCentralTrendsService {
	return &CalculateCentralTrendsService{}
}

func (c *CalculateCentralTrendsService) Calculate(data map[string][]any) (map[string][]any, error) {
	var headers []string

	for header := range data {
		headers = append(headers, header)
	}

	for _, header := range headers {
		if strings.HasSuffix(header, "_type") {
			continue
		}

		dataType := data[header+"_type"][0]

		switch dataType { // Pra cada tipo vai ter um calculo específico
		case "Qualitativa Nominal": // Moda
			var modalData map[string]int = make(map[string]int)

			for _, value := range data[header] {
				modalData[fmt.Sprintf("%v", value)]++
			}

			modalValue := ""
			modalCount := 0

			for value, count := range modalData {
				if count > modalCount {
					modalValue = value
					modalCount = count
				}
			}

			data[header+"_modal_value"] = []any{modalValue}
			data[header+"_modal_count"] = []any{modalCount}
		case "Qualitativa Ordinal": // Moda e Porcentagem
			var modalData map[string]int = make(map[string]int)

			for _, value := range data[header] {
				modalData[fmt.Sprintf("%v", value)]++
			}

			modalValue := ""
			modalCount := 0

			for value, count := range modalData {
				if count > modalCount {
					modalValue = value
					modalCount = count
				}
			}

			data[header+"_modal_value"] = []any{modalValue}
			data[header+"_modal_count"] = []any{modalCount}
			// Posso calcular a porcentagem de cada 1 tbm...
			total := float64(len(data[header]))

			for value, count := range modalData {
				valuePercent := float64(count) / total
				data[header+"_percent_"+value] = []any{valuePercent}
			}
		case "Binária": // Moda e Proporção
			var modalData map[string]int = make(map[string]int)

			for _, value := range data[header] {
				modalData[fmt.Sprintf("%v", value)]++
			}

			modalValue := ""
			modalCount := 0

			for value, count := range modalData {
				if count > modalCount {
					modalValue = value
					modalCount = count
				}
			}

			data[header+"_modal_value"] = []any{modalValue}
			data[header+"_modal_count"] = []any{modalCount}
			// Posso calcular a porcentagem de cada 1 tbm...
			total := float64(len(data[header]))

			for value, count := range modalData {
				valuePercent := float64(count) / total
				data[header+"_percent_"+value] = []any{valuePercent}
			}
		case "Quantitativa Discreta": // Média, Mediana e Moda
			var modalData map[string]int = make(map[string]int)

			for _, value := range data[header] {
				modalData[fmt.Sprintf("%v", value)]++
			}

			modalValue := ""
			modalCount := 0

			for value, count := range modalData {
				if count > modalCount {
					modalValue = value
					modalCount = count
				}
			}

			data[header+"_modal_value"] = []any{modalValue}
			data[header+"_modal_count"] = []any{modalCount}

			var average float64 = 0

			for _, value := range data[header] {
				intValue, ok := value.(int64)
				if !ok {
					fmt.Printf("Error on converting int value in 'Quantitativa Discreta'. header: %s - value: %v\n", header, value)
					return nil, errors.New("error on converting int value")
				}
				average += float64(intValue)
			}

			data[header+"_average"] = []any{average / float64(len(data[header]))}

			var sortedValues []int64

			for _, value := range data[header] {
				intValue, ok := value.(int64)
				if !ok {
					fmt.Printf("Error on converting int value in 'Quantitativa Discreta'. header: %s - value: %v\n", header, value)
					return nil, errors.New("error on converting int value")
				}
				sortedValues = append(sortedValues, intValue)
			}

			// Eu mesmo faço o sort
			sortValues(&sortedValues)

			if len(sortedValues)%2 == 0 {
				median := (sortedValues[len(sortedValues)/2-1] + sortedValues[len(sortedValues)/2]) / 2
				data[header+"_median"] = []any{median}
			} else {
				median := sortedValues[len(sortedValues)/2]
				data[header+"_median"] = []any{median}
			}

		case "Quantitativa Contínua": // Média
			var average float64 = 0

			for _, value := range data[header] {
				floatValue, ok := value.(float64)
				if !ok {
					fmt.Printf("Error on converting float value in 'Quantitativa Contínua'. header: %s - value: %v\n", header, value)
					return nil, errors.New("error on converting float value")
				}
				average += floatValue
			}

			data[header+"_average"] = []any{average / float64(len(data[header]))}
		}

	}

	return data, nil
}

func sortValues(values *[]int64) {
	for i := 0; i < len(*values); i++ {
		for j := i + 1; j < len(*values); j++ {
			if (*values)[i] > (*values)[j] {
				(*values)[i], (*values)[j] = (*values)[j], (*values)[i]
			}
		}
	}
}
