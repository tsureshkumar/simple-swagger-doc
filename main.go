package main

import (
	"fmt"
	"os"

	"github.com/tsureshkumar/simple-swagger-doc/formats"
	"github.com/tsureshkumar/simple-swagger-doc/formats/markdown"
	"github.com/tsureshkumar/simple-swagger-doc/models"
	"github.com/tsureshkumar/simple-swagger-doc/myerr"
)

func main() {

	formats.Register("md", markdown.NewFormatter)

	// FIXME: validate arguments
	doc, err := models.LoadJson(os.Args[1])
	if err != nil {
		panic(err)
	}
	markdown, err := formats.GetService("md")
	if err != nil {
		fmt.Println(myerr.Wrap(err, "service error"))
	}
	err = markdown.Format(doc)
	if err != nil {
		fmt.Println(myerr.Wrap(err, "formatting error"))
	}

}
