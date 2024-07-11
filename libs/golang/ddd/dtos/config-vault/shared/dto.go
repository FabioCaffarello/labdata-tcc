package shareddto

// JobDependenciesDTO represents the data transfer object for job dependencies.
// It includes the service and source details that are dependent on each other.
type JobDependenciesDTO struct {
	Service string `json:"service"` // Service represents the name of the dependent service.
	Source  string `json:"source"`  // Source indicates the origin or source of the dependency.
}

// JobParametersDTO represents the data transfer object for job parameters.
// It includes the parser module that is used to parse the configuration.
type JobParametersDTO struct {
	ParserModule string `json:"parser_module"` // ParserModule specifies the module used for parsing the configuration.
}
