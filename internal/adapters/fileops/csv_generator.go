package fileops

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"strings"
)

type CSVGenerator struct{}

func NewCSVGenerator() *CSVGenerator {
	return &CSVGenerator{}
}

func (c CSVGenerator) GenerateCSV(data map[string][]any) (*bytes.Buffer, error) {

	buf := &bytes.Buffer{}
	writer := csv.NewWriter(buf)

	headers := make([]string, 0)
	for header := range data {
		headers = append(headers, header)
	}

	// Escrevendo Cabecalho
	if err := writer.Write(headers); err != nil {
		return nil, err
	}

	// Escrevendo as linhas
	rowCount := 0
	for header := range data {
		if strings.HasSuffix(header, "_type") {
			continue
		}
		rowCount = len(data[header])
		break
	}

	for i := 0; i < rowCount; i++ {
		row := make([]string, len(headers))
		for j, header := range headers {
			if i >= len(data[header]) {
				row[j] = ""
				continue
			}
			row[j] = fmt.Sprintf("%v", data[header][i])
		}
		if err := writer.Write(row); err != nil {
			return nil, err
		}
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		return nil, err
	}

	return buf, nil
}
