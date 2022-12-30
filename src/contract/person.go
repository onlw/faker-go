package contract

import "github.com/onlw/faker-go/v2/src/model"

type Person interface {
	Name() string

	Person() model.Person
}
