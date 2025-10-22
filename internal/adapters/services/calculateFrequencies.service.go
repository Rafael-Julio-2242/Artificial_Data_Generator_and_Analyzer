package services

import (
	"fmt"
	"sort"
	"strings"
)

type CalculateFrequenciesService struct{}

func NewCalculateFrequenciesService() *CalculateFrequenciesService {
	return &CalculateFrequenciesService{}
}

func (c *CalculateFrequenciesService) Calculate(data map[string][]any) (map[string][]any, error) {

	for header, values := range data {
		if strings.HasSuffix(header, "_type") {
			continue
		}

		var absoluteFrequencies map[string]int = make(map[string]int, 0)

		for _, value := range values {
			absoluteFrequencies[fmt.Sprintf("%v", value)]++
		}

		var relativeFrequencies map[string]float64 = make(map[string]float64, 0)

		for value, count := range absoluteFrequencies {
			relativeFrequencies[value] = float64(count) / float64(len(values))
		}

		// Ordenar as chaves para garantir a ordem correta do acumulado
		sortedValues := make([]string, 0, len(absoluteFrequencies))
		for value := range absoluteFrequencies {
			sortedValues = append(sortedValues, value)
		}
		sort.Strings(sortedValues)

		// Calcular frequências acumuladas
		acumulativeFrequencies := make(map[string]float64, len(absoluteFrequencies))
		acumulativeSum := 0.0

		for _, value := range sortedValues {
			acumulativeSum += float64(absoluteFrequencies[value])
			acumulativeFrequencies[value] = acumulativeSum
		}

		data[header+"_absolute_frequencies"] = []any{absoluteFrequencies}
		data[header+"_relative_frequencies"] = []any{relativeFrequencies}
		data[header+"_acumulative_frequencies"] = []any{acumulativeFrequencies}
	}

	return data, nil
}
