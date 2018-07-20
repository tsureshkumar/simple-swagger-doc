package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"

	"github.com/tsureshkumar/simple-swagger-doc/myerr"
)

type version struct {
	OpenAPI string `json:"openapi"`
	Swagger string `json:"swagger"`
}

func loadContent(filename string) ([]byte, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, myerr.Wrap(err, "file reading failed")
	}
	return content, nil
}

func checkSchemaVersion(content string) (string, string, error) {
	var re = regexp.MustCompile(`.*(?P<model>swagger|openapi)["']\s*:\s*['"](?P<version>[^"']*)['"].*`)
	match := re.FindStringSubmatch(content)
	model := "swagger"
	version := "2.0"
	if len(match) < 3 {
		fmt.Printf("Warning: assuming 2.0 as swagger version")
	} else {
		for i, name := range re.SubexpNames() {
			switch name {
			case "model":
				model = match[i]
			case "version":
				version = match[i]
			}
		}
	}
	return model, version, nil
}

// Loads given json file into Swagger or OpenAPI models. Checks the schema
// version of the file and loads appropriate models for later processing.
// Retruns error if schema is not "swagger: 2.0" or "openapi: 3.0.[01]".
func LoadJson(filename string) (*Document, error) {
	content, err := loadContent(filename)
	if err != nil {
		return nil, err
	}
	model, version, err := checkSchemaVersion(string(content))
	cmodel := CSwagger
	switch model {
	case "swagger":
		cmodel = CSwagger
	case "openapi":
		cmodel = COpenAPI
	}
	doc := &Document{Version: version, Model: cmodel}
	if doc.Model == CSwagger && doc.Version == "2.0" {
		swagger := &Swagger2{}
		if err = json.Unmarshal(content, swagger); err != nil {
			return nil, myerr.Wrap(err, "unmarshalling swagger2 error")
		}
		doc.Swagger2 = swagger
	} else {
		openapi := &OpenAPI{}
		if err = json.Unmarshal(content, openapi); err != nil {
			return nil, myerr.Wrap(err, "unmarshalling openapi error")
		}
		doc.OpenAPI = openapi

	}

	return doc, nil
}
