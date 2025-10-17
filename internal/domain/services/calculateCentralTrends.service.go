package services

type CalculateCentralTrends interface {
	Calculate(data map[string][]any) (map[string][]any, error)
}

const CalculateCentralTrendsServiceKey = "calculateCentralTrendsService"
