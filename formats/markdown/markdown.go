package markdown

import (
	"fmt"
	"os"
	"reflect"
	"regexp"
	"sort"
	"strings"

	"github.com/tsureshkumar/simple-swagger-doc/formats"
	"github.com/tsureshkumar/simple-swagger-doc/models"
	"github.com/tsureshkumar/simple-swagger-doc/myerr"
	"github.com/tsureshkumar/simple-swagger-doc/util"
)

var methods = []string{"Get", "Post"}

type formatter struct {
}

// FIXME: should work more than the definitions/components
func dereferenceLocal(doc *models.Document, ref string) (string, interface{}, error) {
	if ref[0] != '#' {
		return "", nil, myerr.New("local reference shoudl start with #")
	}
	var obj interface{}
	var name string
	switch doc.Model {
	case models.CSwagger:
		obj = doc.Swagger2
	case models.COpenAPI:
		obj = doc.OpenAPI
	}

	paths := strings.Split(ref[1:], "/")
	for _, p := range paths {
		r := reflect.ValueOf(obj)
		if p != "" {
			if r.Kind() == reflect.Ptr || r.Kind() == reflect.Struct {
				f := reflect.Indirect(r).FieldByName(strings.Title(p))
				obj = f.Interface()
			} else {
				switch v := obj.(type) {
				case map[string]*models.Schema:
					obj = v[p]
				}
			}
			name = p
		}
	}
	return name, obj, nil
}

func printHeader(info *models.Info) error {
	fmt.Printf("# %s\n", info.Title)
	fmt.Printf("%s\n\n", info.Description)
	return nil
}

func printShortAPIs(keys []string, paths map[string]models.PathItem, id *int) {
	fmt.Printf("# APIs\n")

	fmt.Printf("| Method | Path | Summary |\n")
	fmt.Printf("|--|--|--|\n")
	for _, pk := range keys {
		pv := paths[pk]
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

func printInputParams(op *models.Operation) {
	if len(op.Parameters) == 0 {
		return
	}
	fmt.Printf("\n### Request Parameters\n")

	headers := []string{"Name", "Location", "Type", "Description"}

	fmt.Printf("\n| %s |\n", strings.Join(headers, " | "))
	fmt.Printf("|%s|\n", strings.Join(util.Map(headers, func(x string) string { return "-" }), "-|-"))
	for _, param := range op.Parameters {
		if param.Ref == nil {
			typeRef := param.Type
			if typeRef == "" && param.Schema != nil {
				typeRef = param.Schema.Type
				if param.Schema.Ref != nil {
					typeRef = fmt.Sprintf("<a href='%s'>%s<a>", param.Schema.Ref.Ref,
						strings.Replace(param.Schema.Ref.Ref, "#/definitions/", "", 1))
				}
				if len(param.Schema.Enum) != 0 {
					typeRef += "<br> [" + strings.Join(param.Schema.Enum, ",") + "]"
				}
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
	if len(op.Responses) == 0 {
		return
	}
	if doc.Model == models.CSwagger {
		headers := []string{"Code", "Data", "Description"}

		fmt.Printf("\n| %s |\n", strings.Join(headers, " | "))
		fmt.Printf("|%s|\n", strings.Join(util.Map(headers, func(x string) string { return "-" }), "-|-"))
		for k, resp := range op.Responses {
			var typeRef = ""
			if resp.Schema != nil {
				if resp.Schema.Ref != nil {
					typeRef = fmt.Sprintf("<a href='%s'>%s<a>", resp.Schema.Ref.Ref, strings.Replace(resp.Schema.Ref.Ref, "#/definitions/", "", 1))
				} else {
					var sb strings.Builder
					printSchema(doc, "", resp.Schema, &sb, 1)
					typeRef = sb.String()
				}
			}
			fmt.Printf("| %s | %s | %s |\n",
				k, typeRef, resp.Description)
		}
	} else {
		printOpenAPIResponses(doc, op)
	}
}

func printOpenAPIResponses(doc *models.Document, op *models.Operation) {
	headers := []string{"Code", "Content Type", "Data", "Description"}
	fmt.Printf("\n| %s |\n", strings.Join(headers, " | "))
	fmt.Printf("|%s|\n", strings.Join(util.Map(headers, func(x string) string { return "-" }), "-|-"))
	for k, resp := range op.Responses {
		if len(resp.Content) == 0 {
			fmt.Printf("| %s | %s | %s | %s |\n", k, "", "", resp.Description)
		}
		for cv, content := range resp.Content {
			var typeRef = ""
			if content.Schema != nil {
				if content.Schema.Ref != nil {
					typeRef = fmt.Sprintf("<a href='%s'>%s<a>", content.Schema.Ref.Ref, strings.Replace(content.Schema.Ref.Ref, "#/components/schemas/", "", 1))
				} else {
					var sb strings.Builder
					printSchema(doc, "", content.Schema, &sb, 1)
					re := regexp.MustCompile(`\r?\n`)
					v := re.ReplaceAllString(sb.String(), "<br>")
					re = regexp.MustCompile(`\s`)
					v = re.ReplaceAllString(v, "&nbsp;")
					typeRef = "<code>" + v + "</code>"
				}
			}
			fmt.Printf("| %s | %s | %s | %s |\n", k, cv, typeRef, resp.Description)
		}
	}
}

func printRequestBody(doc *models.Document, op *models.Operation) {
	if op.RequestBody == nil {
		return
	}
	if doc.Model == models.CSwagger {
	} else {
		fmt.Println("\n### Request Body")
		fmt.Println(op.RequestBody.Description)
		if op.RequestBody.Description != "" {
			fmt.Printf("\n")
		}
		headers := []string{"Content Type", "Data"}
		fmt.Printf("\n| %s |\n", strings.Join(headers, " | "))
		fmt.Printf("|%s|\n", strings.Join(util.Map(headers, func(x string) string { return "-" }), "-|-"))

		for cv, content := range op.RequestBody.Content {
			var typeRef = ""
			if content.Schema != nil && content.Schema.Ref != nil {
				typeRef = fmt.Sprintf("<a href='%s'>%s<a>", content.Schema.Ref.Ref, strings.Replace(content.Schema.Ref.Ref, "#/components/schemas/", "", 1))
			} else {
				// FIXME: print the request object directly
			}

			fmt.Printf("| %s | %s |\n", cv, typeRef)
		}
	}
}

func printDetailedAPIs(doc *models.Document, paths map[string]models.PathItem, id *int) {
	fmt.Printf("# API Details\n")
	keys := make([]string, 1)
	for k := range paths {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, path := range keys {
		pv := paths[path]
		for _, m := range methods {
			f := reflect.ValueOf(pv).FieldByName(m)
			op := f.Interface().(*models.Operation)
			if op != nil {
				opID := op.OperationId
				fmt.Printf("\n## %s %s\n", strings.ToUpper(m), path)
				fmt.Printf("<a name='%s'>%s</a>\n", opID, op.Summary)
				fmt.Printf("\n%s\n", op.Description)
				printInputParams(op)
				printRequestBody(doc, op)
				fmt.Printf("\n### Responses\n")
				printResponses(doc, op)
			}
		}
	}
}

func fixOperationIDs(paths map[string]models.PathItem) {
	id := 0
	for _, pv := range paths {
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

func printSchema(doc *models.Document, name string, sch *models.Schema, sb *strings.Builder, level int) {
	prefix := strings.Repeat(" ", level*2)
	desc := ""
	if sch.Description != "" {
		desc = fmt.Sprintf("\t// %s", sch.Description)
	}
	if sch.AllOf != nil {
		for _, schitems := range sch.AllOf {
			if schitems.Ref != nil {
				refName, refSchema, err := dereferenceLocal(doc, schitems.Ref.Ref)
				if err != nil {
					fmt.Fprintf(os.Stderr, "error referencing local ref '%s'\n", schitems.Ref.Ref)
				} else {
					schema := refSchema.(*models.Schema)
					printSchema(doc, refName, schema, sb, level+1)
				}
			} else {
				printSchema(doc, name, schitems, sb, level+1)
			}
		}
		return
	}
	sb.WriteString(fmt.Sprintf("%s %s %s %s\n", prefix, name, sch.Type, desc))
	for k, v := range sch.Properties {
		printSchema(doc, k, v, sb, level+1)
	}
	if sch.Enum != nil {
		nextLevelPrefix := strings.Repeat(" ", (level+1)*2)
		nextNextLevelPrefix := strings.Repeat(" ", (level+2)*2)
		sb.WriteString(fmt.Sprintf("%s enum:\n", nextLevelPrefix))
		for _, vv := range sch.Enum {
			sb.WriteString(fmt.Sprintf("%s %s\n", nextNextLevelPrefix, vv))
		}
	}
}

func printDefinitions(doc *models.Document, defs map[string]models.Definition) {
	fmt.Println("# Definitions")

	for dk, dv := range defs {
		fmt.Printf("## %s\n", dk)
		fmt.Printf("<a name='/definitions/%s'>type: %s</a>\n\n", dk, dv.Type)
		var sb strings.Builder
		printSchema(doc, dk, dv.Schema, &sb, 1)
		fmt.Printf("```\n%s```\n", sb.String())
	}
}

func printComponents(doc *models.Document, comps models.Components) {
	fmt.Println("# Components")

	fmt.Println("# Schema Objects")
	for dk, dv := range comps.Schemas {
		fmt.Printf("## %s\n", dk)
		if dv.Type == "" {
			dv.Type = "object"
		}
		fmt.Printf("<a name='/components/schemas/%s'>type: %s</a>\n\n", dk, dv.Type)
		var sb strings.Builder
		printSchema(doc, dk, dv, &sb, 1)
		fmt.Printf("```\n%s```\n", sb.String())
	}

}

func sortPaths(paths map[string]models.PathItem) []string {
	keys := make([]string, 1)
	for k := range paths {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// factory implementation. Formats the given document into markdown and prints
// the result to stdout. The processing varies based on the type of document.
func (f formatter) Format(doc *models.Document) error {
	id := 0
	if doc.Model == models.CSwagger {
		fixOperationIDs(doc.Swagger2.Paths)
		printHeader(&doc.Swagger2.Info)
		sortedPaths := sortPaths(doc.Swagger2.Paths)
		printShortAPIs(sortedPaths, doc.Swagger2.Paths, &id)
		printDetailedAPIs(doc, doc.Swagger2.Paths, &id)
		printDefinitions(doc, doc.Swagger2.Definitions)
	} else {
		fixOperationIDs(doc.OpenAPI.Paths)
		printHeader(&doc.OpenAPI.Info)
		sortedPaths := sortPaths(doc.OpenAPI.Paths)
		printShortAPIs(sortedPaths, doc.OpenAPI.Paths, &id)
		printDetailedAPIs(doc, doc.OpenAPI.Paths, &id)
		printComponents(doc, doc.OpenAPI.Components)
	}
	return nil
}

// Creates an object of the plugin
func NewFormatter() (formats.OutputType, error) {
	f := formatter{}
	return f, nil
}
