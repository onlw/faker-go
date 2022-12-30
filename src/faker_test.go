package src

import (
	"github.com/onlw/faker-go/v2/src/consts"
	"github.com/onlw/faker-go/v2/src/faker"
	"testing"
)

func TestPerson(*testing.T) {
	name := faker.NewFaker().Locale(consts.LocaleZhCn).Person().Name()
	println(name)
	name = faker.NewFaker().Locale(consts.LocaleEnUs).Person().Name()
	println(name)
}

func TestAddress(t *testing.T) {
	val := faker.NewFaker().Address().AddressName()
	t.Logf(val)
	val = faker.NewFaker().Locale(consts.LocaleZhCn).Address().AddressName()
	t.Logf(val)
}

//func ExampleAddress() {
//	fmt.Println(faker.NewFaker().Address().AddressName())
//	// Output: 364 Unions ville, Norfolk, Ohio 99536
//}
