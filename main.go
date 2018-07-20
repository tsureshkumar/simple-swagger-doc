package main

import (
	"fmt"
	"os"

	"github.com/tsureshkumar/simple-swagger-doc/formats"
	"github.com/tsureshkumar/simple-swagger-doc/formats/markdown"
	"github.com/tsureshkumar/simple-swagger-doc/models"
	"github.com/tsureshkumar/simple-swagger-doc/myerr"
)

func usage() {
	fmt.Printf("Usage: %s openapi-schema.json\n", os.Args[0])
}

func main() {

	formats.Register("md", markdown.NewFormatter)

	// validate cmd line args
	if len(os.Args) <= 1 {
		usage()
		os.Exit(1)
	}

	// load json file into models
	doc, err := models.LoadJson(os.Args[1])
	if err != nil {
		panic(err)
	}

	// output markdown
	markdown, err := formats.GetService("md")
	if err != nil {
		fmt.Println(myerr.Wrap(err, "service error"))
	}
	err = markdown.Format(doc)
	if err != nil {
		fmt.Println(myerr.Wrap(err, "formatting error"))
	}

}
