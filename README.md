# Simple Swagger API Doc

This project helps to create a simple markdown version from Swagger API
specification. 

Often developers create the API specification using swagger editor and use
swagger UI to host the API in a browsable manner.  But, it is hard to take a PDF
version or integrate with other documentation systems, like confluence. This
project solves that by simply exporting the API into a markdown. From markdown,
it is easier to convert to other formats using other tools.

![Preview](https://github.com/tsureshkumar/my-docs/raw/master/images/simple-swagger-doc-screenshot.png)

# How to run?

You need to have the json version of the API.  You can download your API in json
format using the swagger editor. Once you have it, you can convert into a simple
markdown as 

```
$ go run main.go samples/petstore.json 
```

If you have downloaded the binary, run it as 
```
$ ./simple-swagger-doc samples/petstore-openapi.json
```

# Converting into HTML

You can convert the markdown into html using the pandoc utility.

```
$ go run main.go samples/petstore-openapi.json | \
    pandoc -f markdown -css=github-pandoc.css 
```

# Using with confluence

You can create a "CSS StyleSheet" macro in your page and add the content of
the file github-confluence-css-minifieid.css. The create a HTML macro and add
the content of the generated HTML above.

# Samples

You can refer to the final output of some of the samples

* [Pet Store - Swagger 2.0](https://github.com/tsureshkumar/simple-swagger-doc/blob/master/samples/petstore.md)
* [Pet Store - OpenAPI ](https://github.com/tsureshkumar/simple-swagger-doc/blob/master/samples/petstore-openapi.md)
* [Pet Store Extended - OpenAPI ](https://github.com/tsureshkumar/simple-swagger-doc/blob/master/samples/petstore-openapi-extended.md)

# Supported Swagger API versions

I have tested few samples of swagger schema version 2.0 and openapi 3.0.1

# Limitations

This tool supports most of the basic directives of the Swagger API.  Some of the
things that will not work yet are

* External reference documents for schema/components
* OpenAPI components other than Schema. But, this is trivial to fix

Also, not all content in the API is transfered to the markdown to avoid
cluttering.  The resulting document is very simple and readable for other
developers and architects. This will be gradually enhanced.
