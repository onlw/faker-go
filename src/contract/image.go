package contract

type Image interface {
	ImageURL(width int, height int) string

	//Image() model.Image
}
