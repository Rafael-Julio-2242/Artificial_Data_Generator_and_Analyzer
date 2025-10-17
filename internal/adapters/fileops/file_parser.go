package fileops

import (
	"encoding/csv"
	"errors"
	"io"
	"mime/multipart"
	"strings"

	"github.com/xuri/excelize/v2"
)

type FileParser struct {
}

func NewFileParser() *FileParser {
	return &FileParser{}
}

func (f FileParser) handleCsvData(fileData *multipart.File, fname string) (map[string][]string, error) {

	isCsv := strings.HasSuffix(fname, ".csv")

	if !isCsv {
		return nil, errors.New("mimetype does not correspond to file suffix")
	}

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

func (f FileParser) handleXlsxData(fileData *multipart.File, fname string) (map[string][]string, error) {

	isXlsx := strings.HasSuffix(fname, ".xlsx")

	if !isXlsx {
		return nil, errors.New("mimetype does not correspond to file suffix")
	}

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

func (f FileParser) handleJsonData(fileData *multipart.File, fname string) (map[string][]string, error) {
	return nil, nil // TODO Implement
}

func (f FileParser) handleTsvData(fileData *multipart.File, fname string) (map[string][]string, error) {
	return nil, nil // TODO Implement
}

func (f FileParser) ConvertFileToData(fileData *multipart.File, ext string, fname string) (map[string][]string, error) {
	switch ext {
	case "text/csv":
		return f.handleCsvData(fileData, fname)
	case "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet":
		return f.handleXlsxData(fileData, fname)
	case "application/json":
		return f.handleJsonData(fileData, fname)
	case "text/tab-separated-values":
		return f.handleTsvData(fileData, fname)
	default:
		return nil, errors.New("file type not supported")
	}
}
