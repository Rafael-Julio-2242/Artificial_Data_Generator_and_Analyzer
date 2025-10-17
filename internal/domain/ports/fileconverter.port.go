package ports

import "mime/multipart"

type FileParser interface {
	ConvertFileToData(fileData *multipart.File, ext string, fname string) (map[string][]string, error)
}

const FileParserServiceKey = "fileparser"
