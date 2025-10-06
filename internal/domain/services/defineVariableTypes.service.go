package services

type DefineVariableTypesService interface {
	DefineVariableTypes(data map[string][]string) (map[string][]any, error)
}

const DefineVariableTypesServiceKey = "defineVariableTypesService"
