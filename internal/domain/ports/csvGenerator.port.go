package ports

import "bytes"

type CSVGenerator interface {
	GenerateCSV(data map[string][]any) (*bytes.Buffer, error)
}

const CSVGeneratorServiceKey = "csvGenerator"
