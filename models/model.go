package models

type ModelType int

const (
	CSwagger ModelType = iota + 1
	COpenAPI
)

type Document struct {
	Swagger2 *Swagger2
	OpenAPI  *OpenAPI
	Model    ModelType
	Version  string
}
