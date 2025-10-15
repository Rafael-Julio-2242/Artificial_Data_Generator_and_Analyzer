package services

import (
	"context"
	"fmt"
	"strconv"

	"google.golang.org/genai"
)

type DefineVariableTypesService struct {
	aiClient *genai.Client
	ctx      context.Context
}

func NewDefineVariableTypesService(aiClient *genai.Client, ctx context.Context) *DefineVariableTypesService {

	return &DefineVariableTypesService{
		aiClient: aiClient,
		ctx:      ctx,
	}
}

func (s *DefineVariableTypesService) DefineVariableTypes(data map[string][]string) (map[string][]any, error) {
	// TODO Falta apenas validar se a classificação é binária
	fixedData, err := s.fixDataTyping(data)
	resultData := make(map[string][]any, len(fixedData))

	if err != nil {
		return nil, err
	}

	for header, values := range fixedData {
		resultData[header] = values
	}

	for header, values := range fixedData {

		value := values[0]

		switch any(value).(type) {
		case int64:
			keyValue := fmt.Sprintf("%v_type", header)

			// Antes de definir como discreta, existe a possibilidade de ser binária.
			// Tenho que validar valor a valor pra descobrir isso

			value1 := values[0]
			var value2 int64
			isBinary := true

			for _, value := range values {
				intValue, ok := value.(int64)

				if !ok {
					continue
				}

				if value1 != intValue && value2 == 0 {
					value2 = intValue
					continue
				}

				if value2 != 0 && value1 != intValue && value2 != intValue {
					isBinary = false
					break
				}

			}

			if isBinary {
				value := []any{"Binária"}
				resultData[keyValue] = value
			} else {
				value := []any{"Quantitativa Discreta"}
				resultData[keyValue] = value
			}
		case float64:
			keyValue := fmt.Sprintf("%v_type", header)
			value := []any{"Quantitativa Contínua"}
			resultData[keyValue] = value
		case string:
			keyValue := fmt.Sprintf("%v_type", header)

			aiClassificationResponse, err := s.aiClient.Models.GenerateContent(
				s.ctx,
				"gemini-2.5-flash",
				genai.Text("Que tipo de dado é esse? Responda apenas com uma das opções: Qualitativa Nominal ou Qualitativa Ordinal ?\n Dado: "+value.(string)),
				nil,
			)

			if err != nil {
				return nil, err
			}

			textResponse := aiClassificationResponse.Text()

			value := []any{textResponse}
			resultData[keyValue] = value
		}

	}

	return resultData, nil
}

func (s *DefineVariableTypesService) fixDataTyping(data map[string][]string) (map[string][]any, error) {
	resultDataSet := make(map[string][]any, len(data))

	for header, values := range data {

		value := values[0]

		_, err := strconv.ParseInt(value, 10, 64)

		if err == nil {
			// is int
			parsed := make([]any, len(values))
			for index, value := range values {
				intValue, err := strconv.ParseInt(value, 10, 64)
				var anyValue any = intValue
				if err != nil {
					anyValue = ""
				}
				parsed[index] = anyValue
			}
			resultDataSet[header] = parsed
			continue
		}

		_, err = strconv.ParseFloat(value, 64)

		if err == nil {
			// is float

			parsed := make([]any, len(values))
			for index, value := range values {
				floatValue, err := strconv.ParseFloat(value, 64)
				var anyValue any = floatValue
				if err != nil {
					anyValue = ""
				}
				parsed[index] = anyValue
			}
			resultDataSet[header] = parsed
			continue
		}

		parsed := make([]any, len(values))
		for index, value := range values {
			parsed[index] = value
		}
		// is string
		resultDataSet[header] = parsed

	}

	return resultDataSet, nil
}
