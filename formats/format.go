package formats

import (
	"fmt"

	"github.com/tsureshkumar/simple-swagger-doc/models"
	"github.com/tsureshkumar/simple-swagger-doc/myerr"
)

// Plugin interface for output formats.  Different supported plugins like
// markdown, html ,etc should implement this interface
type OutputType interface {
	Format(*models.Document) error
}

// Factory method to create objects of OutputType plugin
type OutputTypeFactory func() (OutputType, error)

var factory = make(map[string]OutputTypeFactory)

// Register a factory method with name for output plugins
func Register(name string, factoryFn OutputTypeFactory) error {
	_, ok := factory[name]
	if ok {
		return myerr.New("factory already registered")
	}
	factory[name] = factoryFn
	return nil
}

// Utility method to get a plugin given the plugin name
func GetService(name string) (OutputType, error) {
	fact, ok := factory[name]
	if !ok {
		return nil, myerr.New(fmt.Sprintf("no formatter available for '%s'", name))
	}
	svc, err := fact()
	return svc, err
}
