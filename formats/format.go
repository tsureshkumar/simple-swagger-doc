package formats

import (
	"fmt"

	"github.com/tsureshkumar/simple-swagger-doc/models"
	"github.com/tsureshkumar/simple-swagger-doc/myerr"
)

type OutputType interface {
	Format(*models.Document) error
}

// factory
type OutputTypeFactory func() (OutputType, error)

var factory = make(map[string]OutputTypeFactory)

func Register(name string, factoryFn OutputTypeFactory) error {
	_, ok := factory[name]
	if ok {
		return myerr.New("factory already registered")
	}
	factory[name] = factoryFn
	return nil
}

func GetService(name string) (OutputType, error) {
	fact, ok := factory[name]
	if !ok {
		return nil, myerr.New(fmt.Sprintf("no formatter available for '%s'", name))
	}
	svc, err := fact()
	return svc, err
}
