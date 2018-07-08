package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"

	"github.com/au5ton/go-stego/encoder"
)

type Image interface {
	image.Image
	Set(x, y int, c color.Color)
}

var input = flag.String("i", "", "Input file name of PNG")
var output = flag.String("o", "", "Output file name of PNG")

func main() {
	flag.Parse()

	encoder.DoThing()

	if *input == "" {
		log.Fatal("input cannot be empty")
	}
	if *output == "" {
		log.Fatal("output cannot be empty")
	}

	reader, err := os.Open(*input)
	if err != nil {
		log.Fatal(err)
	}

	img, err := png.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	dimg, ok := img.(draw.Image)
	if !ok {
		fmt.Println("img is not a drawable image")
	}
	//i := 0
	for x := 0; x < dimg.Bounds().Max.X; x++ {
		for y := 0; y < dimg.Bounds().Max.Y; y++ {
			r, g, b, _ := dimg.At(x, y).RGBA()

			dimg.Set(x, y, color.RGBA{uint8(255 - int(r/257)), uint8(255 - int(g/257)), uint8(255 - int(b/257)), 255})
			//fmt.Printf("%d %d %d @ %d\n", int(r/257), int(g/257), int(b/257), x)
			//fmt.Println(r)
			//i++
		}
	}

	f, err := os.Create(*output)
	if err != nil {
		log.Fatal(err)
	}
	if err := png.Encode(f, dimg); err != nil {
		f.Close()
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("min x,y: %d,%d\n", img.Bounds().Min.X, img.Bounds().Min.Y)
	fmt.Printf("max x,y: %d,%d\n", img.Bounds().Max.X, img.Bounds().Max.Y)

}
