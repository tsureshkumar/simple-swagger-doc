package markdown

import (
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/tsureshkumar/simple-swagger-doc/formats"
	"github.com/tsureshkumar/simple-swagger-doc/models"
	"github.com/tsureshkumar/simple-swagger-doc/util"
)

var methods = []string{"Get", "Post"}

type formatter struct {
}

func printHeader(doc *models.Document) error {
	fmt.Printf("# %s\n", doc.Swagger2.Info.Title)
	fmt.Printf("%s\n\n", doc.Swagger2.Info.Description)
	return nil
}

func printOperation(op *models.Operation, path string) {
	fmt.Println("GET", path)
}

func printShortAPIs(doc *models.Document, id *int) {
	fmt.Printf("# APIs\n")

	fmt.Printf("| Method | Path | Summary |\n")
	fmt.Printf("|--|--|--|\n")
	keys := make([]string, 1)
	for k := range doc.Swagger2.Paths {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, pk := range keys {
		pv := doc.Swagger2.Paths[pk]
		for _, m := range methods {
			f := reflect.ValueOf(pv).FieldByName(m)
			op := f.Interface().(*models.Operation)
			if op != nil {
				opID := op.OperationId
				fmt.Printf("| %s | %s | %s |\n", strings.ToUpper(m),
					fmt.Sprintf("<a href='#%s'>%s</a>", opID, pk),
					op.Summary)
			}
		}
	}
}

func printInputParams(doc *models.Document, op *models.Operation) {
	fmt.Printf("### Request Parameters\n")

	headers := []string{"Name", "Location", "Type", "Description"}

	fmt.Printf("\n| %s |\n", strings.Join(headers, " | "))
	fmt.Printf("|%s|\n", strings.Join(util.Map(headers, func(x string) string { return "-" }), "-|-"))
	for _, param := range op.Parameters {
		if param.Ref == nil {
			typeRef := param.Type
			if typeRef == "" && param.Schema != nil &&
				param.Schema.Ref != nil {
				typeRef = fmt.Sprintf("<a href='%s'>%s<a>", param.Schema.Ref.Ref, strings.Replace(param.Schema.Ref.Ref, "#/definitions/", "", 1))
			}
			fmt.Printf("| %s | %s | %s | %s |\n",
				param.Name,
				param.In,
				typeRef,
				param.Description)
		}
	}
}

func printResponses(doc *models.Document, op *models.Operation) {
	headers := []string{"Code", "Data", "Description"}

	fmt.Printf("\n| %s |\n", strings.Join(headers, " | "))
	fmt.Printf("|%s|\n", strings.Join(util.Map(headers, func(x string) string { return "-" }), "-|-"))
	for k, resp := range op.Responses {
		var typeRef = ""
		if resp.Schema != nil && resp.Schema.Ref != nil {
			typeRef = fmt.Sprintf("<a href='%s'>%s<a>", resp.Schema.Ref.Ref, strings.Replace(resp.Schema.Ref.Ref, "#/definitions/", "", 1))
		}
		fmt.Printf("| %s | %s | %s |\n",
			k, typeRef, resp.Description)
		if len(resp.Headers) > 0 {
			fmt.Printf("#### Headers\n")
		}
	}
}

func printDetailedAPIs(doc *models.Document, id *int) {
	fmt.Printf("# API Details\n")
	keys := make([]string, 1)
	for k := range doc.Swagger2.Paths {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, path := range keys {
		pv := doc.Swagger2.Paths[path]
		for _, m := range methods {
			f := reflect.ValueOf(pv).FieldByName(m)
			op := f.Interface().(*models.Operation)
			if op != nil {
				opID := op.OperationId
				fmt.Printf("## %s %s\n", strings.ToUpper(m), path)
				fmt.Printf("<a name='%s'>%s</a>\n", opID, op.Summary)
				fmt.Printf("%s\n", op.Description)
				printInputParams(doc, op)
				fmt.Printf("### Responses\n")
				printResponses(doc, op)
			}
		}
	}
}

func fixOperationIDs(doc *models.Document) {
	id := 0
	for _, pv := range doc.Swagger2.Paths {
		for _, m := range methods {
			f := reflect.ValueOf(pv).FieldByName(m)
			op := f.Interface().(*models.Operation)
			if op != nil {
				opID := op.OperationId
				if opID == "" {
					op.OperationId = fmt.Sprintf("%s_%d", m, id)
					id++
				}
			}
		}
	}
}

func printSchema(name string, sch *models.Schema, sb *strings.Builder, level int) {
	prefix := strings.Repeat(" ", level*2)
	desc := ""
	if sch.Description != "" {
		desc = fmt.Sprintf("\t// %s", sch.Description)
	}
	sb.WriteString(fmt.Sprintf("%s %s %s %s\n", prefix, name, sch.Type, desc))
	if sch.Type == "object" {
		for k, v := range sch.Properties {
			printSchema(k, v, sb, level+1)
		}
	} else if sch.Enum != nil {
		nextLevelPrefix := strings.Repeat(" ", (level+1)*2)
		nextNextLevelPrefix := strings.Repeat(" ", (level+2)*2)
		sb.WriteString(fmt.Sprintf("%s enum:\n", nextLevelPrefix))
		for _, vv := range sch.Enum {
			sb.WriteString(fmt.Sprintf("%s %s\n", nextNextLevelPrefix, vv))
		}
	}
}

func printDefinitions(doc *models.Document) {
	fmt.Println("# Definitions")

	for dk, dv := range doc.Swagger2.Definitions {
		fmt.Printf("## %s\n", dk)
		fmt.Printf("<a name='/definitions/%s'>type: %s</a>\n\n", dk, dv.Type)
		var sb strings.Builder
		printSchema(dk, dv.Schema, &sb, 1)
		fmt.Printf("```\n%s```\n", sb.String())
	}
}

func (f formatter) Format(doc *models.Document) error {
	//str, _ := json.Marshal(doc)
	//fmt.Printf("%s\n", string(str))
	id := 0
	fixOperationIDs(doc)
	printHeader(doc)
	printShortAPIs(doc, &id)
	printDetailedAPIs(doc, &id)
	printDefinitions(doc)
	return nil
}

func NewFormatter() (formats.OutputType, error) {
	f := formatter{}
	return f, nil
}
