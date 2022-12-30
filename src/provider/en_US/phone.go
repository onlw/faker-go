package en_US

import "github.com/onlw/faker-go/v2/src/util"

type Phone struct {
}

var data = map[string][]string{
	"prefix": {"1"},
	"suffix": {"###-#######"},
}

func (p Phone) PhoneNumber() string {
	return util.GetRandValue(data["prefix"]) + "-" + util.ReplaceWithNumbers(util.GetRandValue(data["suffix"]))
}
