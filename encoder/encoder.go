package encoder

import (
	"image"
	"image/color"
	"image/draw"
)

type Image interface {
	image.Image
	Set(x, y int, c color.Color)
}

func indexToCoordinate(rect image.Rectangle, index int) (p image.Point) {
	p.X = index % rect.Max.X
	p.Y = int(index / rect.Max.X)
	return
}

func coordinateToIndex(rect image.Rectangle, p image.Point) (index int) {
	index = rect.Max.X*p.Y + rect.Max.Y
	return
}

func EncodeDataFromImage(img draw.Image, data []byte, dataType string) image.Image {
	for i := 0; i < len(data); i++ {
		coord := indexToCoordinate(img.Bounds(), i)
		img.Set(coord.X, coord.Y, encodeByteInPixel(img.At(coord.X, coord.Y).RGBA(), data[i]))
	}
}

func encodeByteInPixel(pix color.RGBA, b byte) (enc color.RGBA) {
	enc = color.RGBA{pix.R, pix.G, pix.B, pix.A}
	return
}
