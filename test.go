package main

import (
	"fmt"
	"image"

	"github.com/au5ton/go-stego/encoder"
)

func main() {
	rect := image.Rectangle{image.Point{0, 0}, image.Point{5, 4}}

	p := image.Point{3, 1}
	fmt.Println(encoder.CoordinateToIndex(rect, p))

	for i := 0; i < rect.Max.X*rect.Max.Y; i++ {
		fmt.Printf("%d : ", i)
		fmt.Print(encoder.IndexToCoordinate(rect, i))
		fmt.Printf("\n")
	}

}
