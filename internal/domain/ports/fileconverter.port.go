package ports

import "mime/multipart"

type FileParser interface {
	ConvertFileToData(fileData *multipart.File, ext string) (map[string][]string, error)
}

const FileParserServiceKey = "fileparser"
