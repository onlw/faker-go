package faker

import (
	"github.com/onlw/faker-go/v2/src/consts"
	"testing"
)

func TestPerson(*testing.T) {
	name := NewFaker().Locale(consts.LocaleZhCn).Person().Name()
	println(name)
	name = NewFaker().Locale(consts.LocaleEnUs).Person().Name()
	println(name)
}

func TestAddress(t *testing.T) {
	val := NewFaker().Address().AddressName()
	t.Logf(val)
	val = NewFaker().Locale(consts.LocaleZhCn).Address().AddressName()
	t.Logf(val)
}

func TestPhone(t *testing.T) {
	//val := NewFaker().Phone().PhoneNumber()
	//t.Logf(val)
	val := NewFaker().Locale(consts.LocaleZhCn).Phone().PhoneNumber()
	t.Logf(val)
	val = NewFaker().Locale(consts.LocaleEnUs).Phone().PhoneNumber()
	t.Logf(val)
}

//func ExampleAddress() {
//	fmt.Println(NewFaker().Address().AddressName())
//	// Output: 364 Unions ville, Norfolk, Ohio 99536
//}
