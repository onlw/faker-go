package zh_CN

import "github.com/onlw/faker-go/v2/src/util"

type Phone struct {
}

var data = map[string][]string{
	"prefix": {"134", "135", "136", "137", "138", "139", "147", "150", "151", "152", "157", "158", "159", "178", "182", "183", "184", "187", "188", "130", "131", "132", "145", "155", "156", "176", "185", "186", "133", "153", "177", "180", "181", "189", "170", "171"},
	"suffix": {"########"},
}

func (p Phone) PhoneNumber() string {
	return util.GetRandValue(data["prefix"]) + util.ReplaceWithNumbers(util.GetRandValue(data["suffix"]))
}
