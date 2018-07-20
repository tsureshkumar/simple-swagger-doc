package models

type OpenAPI struct {
	Openapi    string              `json:"openapi"`
	Host       string              `json:"host,omitempty"`
	BasePath   string              `json:"basePath,omitempty"`
	Schemes    []string            `json:"schemes,omitempty"`
	Consumes   []string            `json:"consumes,omitempty"`
	Produces   []string            `json:"produces,omitempty"`
	Info       Info                `json:"info,omitempty"`
	Paths      map[string]PathItem `json:"paths,omitempty"`
	Components Components          `json:"components"`
}

type Components struct {
	Schemas map[string]*Schema `json:"schemas"`
	//Responses       string `json:"responses"`
	//Parameters      string `json:"parameters"`
	//Examples        string `json:"examples"`
	//RequestBodies   string `json:"requestBodies"`
	//Headers         string `json:"headers"`
	//SecuritySchemes string `json:"securitySchemes"`
	//Links           string `json:"links"`
	//Callbacks       string `json:"callbacks"`
}

//openapi
type RequestBody struct {
	Ref
	Description string                `json:"description"`
	Content     map[string]*MediaType `json:"content"`
	Required    bool                  `json:"required"`
}

// openapi
type Server struct {
	Url         string                     `json:"url"`
	Description string                     `json:"description"`
	Variables   map[string]*ServerVariable `json:"variables"`
}

// openapi
type ServerVariable struct {
	Enum        []string `json:"enum"`
	Default     string   `json:"default"`
	Description string   `json:"description"`
}

// openapi
type Conent struct {
	Items map[string]*MediaType `json:"content,omitempty"`
}

// openapi
type MediaType struct {
	Schema *Schema `json:"schema"`
	//Example  string  `json:"example"`
	Examples map[string]Example `json:"examples"`
	Encoding string             `json:"encoding"`
}

// openapi
type Example struct {
	Ref
	Summary       string `json:"summary"`
	Description   string `json:"description"`
	Value         string `json:"value"`
	ExternalValue string `json:"externalValue"`
}
