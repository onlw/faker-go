package src

import (
	"github.com/onlw/faker-go/v2/src/consts"
	"github.com/onlw/faker-go/v2/src/faker"
	"github.com/onlw/faker-go/v2/src/provider/en_US"
	"testing"
)

func TestCreate(*testing.T) {
	//name := zh_CN.Person{}.Name()
	name := en_US.Person{}.Name()
	println(name)
}

func TestPerson(*testing.T) {
	name := faker.NewFaker().Locale(consts.LocaleZhCn).Person().Name()
	println(name)
	name = faker.NewFaker().Locale(consts.LocaleEnUs).Person().Name()
	println(name)
}

func TestAddress(t *testing.T) {
	name := faker.NewFaker().Locale(consts.LocaleZhCn).Address().AddressName()
	println(name)
}
