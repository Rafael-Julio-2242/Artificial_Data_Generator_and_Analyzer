package fileops

import (
	"encoding/csv"
	"errors"
	"io"
	"mime/multipart"

	"github.com/xuri/excelize/v2"
)

type FileParser struct {
}

func NewFileParser() *FileParser {
	return &FileParser{}
}

func (f FileParser) handleCsvData(fileData *multipart.File) (map[string][]string, error) {
	dataMap := make(map[string][]string)

	reader := csv.NewReader(*fileData)
	headers := []string{}
	lineIndex := 0

	for {
		records, err := reader.Read()

		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, errors.New("error reading records in csv file: " + err.Error())
		}

		for index, record := range records {
			if lineIndex == 0 {
				headers = append(headers, record)
				continue
			}

			header := headers[index]
			dataMap[header] = append(dataMap[header], record)
		}
		lineIndex++
	}

	return dataMap, nil
}

func (f FileParser) handleXlsxData(fileData *multipart.File) (map[string][]string, error) {

	reader, err := excelize.OpenReader(*fileData)
	if err != nil {
		return nil, errors.New("error opening xlsx file: " + err.Error())
	}

	sheetName := reader.GetSheetName(1)

	if sheetName == "" {
		return nil, errors.New("sheet not found")
	}

	rows, err := reader.Rows(sheetName)

	if err != nil {
		return nil, errors.New("error reading rows in xlsx file: " + err.Error())
	}

	dataMap := make(map[string][]string)
	headers := []string{}
	lineIndex := 0

	for rows.Next() {
		records, err := rows.Columns()

		if err != nil {
			return nil, errors.New("error reading columns in xlsx file: " + err.Error())
		}

		for index, record := range records {
			if lineIndex == 0 {
				headers = append(headers, record)
				continue
			}

			header := headers[index]
			dataMap[header] = append(dataMap[header], record)
		}

		lineIndex++
	}

	return dataMap, nil
}

func (f FileParser) handleJsonData(fileData *multipart.File) (map[string][]string, error) {
	return nil, nil // TODO Implement
}

func (f FileParser) handleTsvData(fileData *multipart.File) (map[string][]string, error) {
	return nil, nil // TODO Implement
}

func (f FileParser) ConvertFileToData(fileData *multipart.File, ext string) (map[string][]string, error) {

	switch ext {
	case "csv":
		return f.handleCsvData(fileData)
	case "xlsx":
		return f.handleXlsxData(fileData)
	case "json":
		return f.handleJsonData(fileData)
	case "tsv":
		return f.handleTsvData(fileData)
	default:
		return nil, errors.New("unsupported file")
	}
}
