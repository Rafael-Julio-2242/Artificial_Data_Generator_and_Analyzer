package services

type CalculateFrequencies interface {
	Calculate(data map[string][]any) (map[string][]any, error)
}

const CalculateFrequenciesServiceKey = "calculateFrequenciesService"
