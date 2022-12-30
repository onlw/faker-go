package src

import (
	"github.com/onlw/faker-go/v2/src/provider/en_US"
	"testing"
)

func TestCreate(*testing.T) {
	//name := zh_CN.Person{}.Name()
	name := en_US.Person{}.Name()
	println(name)
}
