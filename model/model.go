package model

type LmakerInput struct {
	DtoSqlTag     string
	JSONOmitEmpty bool

	// for create Create cmd
	ServiceName string
	ProjectName string
	ProjectType string
	Hello       string
}
