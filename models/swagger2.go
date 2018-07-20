package models

type Swagger2 struct {
	Swagger     string                `json:"swagger,omitempty"`
	Host        string                `json:"host,omitempty"`
	BasePath    string                `json:"basePath,omitempty"`
	Schemes     []string              `json:"schemes,omitempty"`
	Consumes    []string              `json:"consumes,omitempty"`
	Produces    []string              `json:"produces,omitempty"`
	Info        Info                  `json:"info,omitempty"`
	Paths       map[string]PathItem   `json:"paths,omitempty"`
	Definitions map[string]Definition `json:"definitions,omitempty"`
}

type Info struct {
	Title       string  `json:"title,omitempty"`
	Description string  `json:"description,omitempty"`
	Contact     Contact `json:"contact,omitempty"`
	License     License `json:"license,omitempty"`
	Version     string  `json:"version,omitempty"`
}

type Contact struct {
	Name  string `json:"name,omitempty"`
	Url   string `json:"url,omitempty"`
	Email string `json:"email,omitempty"`
}

type License struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}

type PathItem struct {
	Ref
	Get        *Operation `json:"get,omitempty"`
	Put        *Operation `json:"put,omitempty"`
	Post       *Operation `json:"post,omitempty"`
	Delete     *Operation `json:"delete,omitempty"`
	Options    *Operation `json:"options,omitempty"`
	Head       *Operation `json:"head,omitempty"`
	Patch      *Operation `json:"patch,omitempty"`
	Parameters *Operation `json:"parameters,omitempty"`
}

type Operation struct {
	Tags        []string `json:"tags,omitempty"`
	Summary     string   `json:"summary,omitempty"`
	Description string   `json:"description,omitempty"`
	//ExternalDocs string `json:"externalDocs,omitempty"`
	OperationId string              `json:"operationId,omitempty"`
	Consumes    []string            `json:"consumes,omitempty"`
	Produces    []string            `json:"produces,omitempty"`
	Parameters  []Parameter         `json:"parameters,omitempty"`
	Responses   map[string]Response `json:"responses,omitempty"`
	Schemes     []string            `json:"schemes,omitempty"`
	Deprecated  bool                `json:"deprecated,omitempty"`
	//Security   string   `json:"security,omitempty"`

	//openapi
	RequestBody *RequestBody `json:"requestBody"`
	Callbacks   string       `json:"callbacks"`
	Servers     []Server     `json:"servers"`
}

type Ref struct {
	Ref string `json:"$ref,omitempty"`
}

// FIXME : add additional parameters
type Parameter struct {
	*Ref
	Name             string  `json:"name,omitempty"`
	In               string  `json:"in,omitempty"`
	Description      string  `json:"description,omitempty"`
	Required         bool    `json:"required,omitempty"`
	Schema           *Schema `json:"schema,omitempty"`
	Type             string  `json:"type,omitempty"`
	Format           string  `json:"format,omitempty"`
	AllowEmptyValue  string  `json:"allowEmptyValue,omitempty"`
	Items            Items   `json:"items,omitempty"`
	CollectionFormat string  `json:"collectionFormat,omitempty"`
	Default          string  `json:"default,omitempty"`
}

type Items struct {
	Type             string        `json:"type,omitempty"`
	Format           string        `json:"format,omitempty"`
	Items            *Items        `json:"items,omitempty"`
	CollectionFormat string        `json:"collectionFormat,omitempty"`
	Default          string        `json:"default,omitempty"`
	Maximum          string        `json:"maximum,omitempty"`
	ExclusiveMaximum string        `json:"exclusiveMaximum,omitempty"`
	Minimum          string        `json:"minimum,omitempty"`
	ExclusiveMinimum string        `json:"exclusiveMinimum,omitempty"`
	MaxLength        string        `json:"maxLength,omitempty"`
	MinLength        string        `json:"minLength,omitempty"`
	Pattern          string        `json:"pattern,omitempty"`
	MaxItems         string        `json:"maxItems,omitempty"`
	MinItems         string        `json:"minItems,omitempty"`
	UniqueItems      string        `json:"uniqueItems,omitempty"`
	Enum             []interface{} `json:"enum,omitempty"`
	MultipleOf       string        `json:"multipleOf,omitempty"`
}

// FIXME: implement schema object from json schema
type Schema struct {
	*Ref
	Format      string `json:"format"`
	Title       string `json:"title"`
	Description string `json:"description"`
	//Default          string             `json:"default"`
	MultipleOf       string             `json:"multipleOf"`
	Maximum          string             `json:"maximum"`
	ExclusiveMaximum string             `json:"exclusiveMaximum"`
	Minimum          string             `json:"minimum"`
	ExclusiveMinimum string             `json:"exclusiveMinimum"`
	MaxLength        string             `json:"maxLength"`
	MinLength        string             `json:"minLength"`
	Pattern          string             `json:"pattern"`
	MaxItems         string             `json:"maxItems"`
	MinItems         string             `json:"minItems"`
	UniqueItems      string             `json:"uniqueItems"`
	MaxProperties    string             `json:"maxProperties"`
	MinProperties    string             `json:"minProperties"`
	Required         []string           `json:"required"`
	Enum             []string           `json:"enum"`
	Type             string             `json:"type"`
	Items            Items              `json:"items"`
	AllOf            []*Schema          `json:"allOf"`
	Properties       map[string]*Schema `json:"properties"`
	//AdditionalProperties string             `json:"additionalProperties"`
	Discriminator string `json:"discriminator"`
	ReadOnly      string `json:"readOnly"`
	//Xml           string `json:"xml"`
	ExternalDocs string `json:"externalDocs"`
	Example      string `json:"example"`
}

type Response struct {
	Ref
	Description string                 `json:"description,omitempty"`
	Schema      *Schema                `json:"schema,omitempty"`
	Headers     map[string]Header      `json:"headers,omitempty"`
	Examples    map[string]interface{} `json:"examples,omitempty"`

	// openapi
	Content map[string]*MediaType `json:"content,omitempty"`
}

type Header struct {
	Description      string        `json:"description,omitempty"`
	Type             string        `json:"type,omitempty"`
	Format           string        `json:"format,omitempty"`
	Items            Items         `json:"items,omitempty"`
	CollectionFormat string        `json:"collectionFormat,omitempty"`
	Csv              string        `json:"csv,omitempty"`
	Ssv              string        `json:"ssv,omitempty"`
	Tsv              string        `json:"tsv,omitempty"`
	Pipes            string        `json:"pipes,omitempty"`
	Default          interface{}   `json:"default,omitempty"`
	Maximum          string        `json:"maximum,omitempty"`
	ExclusiveMaximum string        `json:"exclusiveMaximum,omitempty"`
	Minimum          string        `json:"minimum,omitempty"`
	ExclusiveMinimum string        `json:"exclusiveMinimum,omitempty"`
	MaxLength        string        `json:"maxLength,omitempty"`
	MinLength        string        `json:"minLength,omitempty"`
	Pattern          string        `json:"pattern,omitempty"`
	MaxItems         string        `json:"maxItems,omitempty"`
	MinItems         string        `json:"minItems,omitempty"`
	UniqueItems      string        `json:"uniqueItems,omitempty"`
	Enum             []interface{} `json:"enum,omitempty"`
	MultipleOf       string        `json:"multipleOf,omitempty"`
}

type Definition struct {
	*Schema
}
