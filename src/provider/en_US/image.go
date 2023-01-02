package en_US

import (
	"strconv"
)

type Image struct {
}

var imageDomain = "https://picsum.photos/"

func (i Image) ImageURL(width int, height int) string {
	return imageDomain + strconv.Itoa(width) + "/" + strconv.Itoa(height)
}

//func (p Image) Image() model.Image {
//	return model.Image{
//		Url: p.ImageURL(200, 200),
//	}
//}

//func (p Image) name() {
//
//}
