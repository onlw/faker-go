package faker

import (
	"github.com/onlw/faker-go/v2/src/consts"
	"github.com/onlw/faker-go/v2/src/contract"
	"github.com/onlw/faker-go/v2/src/provider/en_US"
	"github.com/onlw/faker-go/v2/src/provider/zh_CN"
	"sync"
)

const defaultLocale = consts.LocaleEnUs

var once sync.Once

var instance *Faker

type Faker struct {
	localeName        string
	personStrategies  map[string]contract.Person
	addressStrategies map[string]contract.Address
	phoneStrategies   map[string]contract.Phone
	companyStrategies map[string]contract.Company
	imageStrategies   map[string]contract.Image
}

var phoneStrategies = map[string]contract.Phone{
	consts.LocaleZhCn: &zh_CN.Phone{},
	consts.LocaleEnUs: &en_US.Phone{},
}

var companyStrategies = map[string]contract.Company{
	consts.LocaleZhCn: &zh_CN.Company{},
	consts.LocaleEnUs: &en_US.Company{},
}
var imageStrategies = map[string]contract.Image{
	consts.LocaleEnUs: &en_US.Image{},
}

func NewFaker() *Faker {
	once.Do(func() {
		instance = &Faker{
			localeName: defaultLocale,
			personStrategies: map[string]contract.Person{
				consts.LocaleZhCn: &zh_CN.Person{},
				consts.LocaleEnUs: &en_US.Person{},
			},
			addressStrategies: map[string]contract.Address{
				consts.LocaleZhCn: &zh_CN.Address{},
				consts.LocaleEnUs: &en_US.Address{},
			},
			phoneStrategies:   phoneStrategies,
			companyStrategies: companyStrategies,
			imageStrategies:   imageStrategies,
		}
	})
	return instance
}

func (f Faker) Locale(locale string) Faker {
	f.localeName = locale

	return f
}

//
//func (f Faker) RegisterStrategy(locale string, strategy interface{}) {
//
//}

func (f Faker) Person() contract.Person {
	return f.personStrategies[f.localeName]
}

func (f Faker) Address() contract.Address {
	return f.addressStrategies[f.localeName]
}

func (f Faker) Phone() contract.Phone {
	return f.phoneStrategies[f.localeName]
}

func (f Faker) Company() contract.Company {
	return f.companyStrategies[f.localeName]
}

func (f Faker) Image() contract.Image {
	return f.imageStrategies[f.localeName]
}

func Locales(locale string) *Faker {
	instance.localeName = locale

	return instance
}

func Image() contract.Image {
	println(instance.imageStrategies)
	return instance.imageStrategies[instance.localeName]
}

//func ImageModel() model.Image {
//	//return instance.imageStrategies[instance.localeName].()
//}
