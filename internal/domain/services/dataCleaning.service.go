package services

type DataCleaningService interface {
	CleanData(data map[string][]string) (map[string][]string, error)
}

const DataCleaningServiceKey = "dataCleaningService"
