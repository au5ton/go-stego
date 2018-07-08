package main

import (
	"flag"
	"fmt"
	"image/png"
	"log"
	"os"
)

var input = flag.String("i", "", "Input file name of PNG")

//var output = flag.String("o", "", "Output file name of PNG")

func main() {
	flag.Parse()

	if *input == "" {
		log.Fatal("input cannot be empty")
	}
	// if *output == "" {
	// 	log.Fatal("output cannot be empty")
	// }

	reader, err := os.Open(*input)
	if err != nil {
		log.Fatal(err)
	}

	img, err := png.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	for x := 0; x < img.Bounds().Max.X; x++ {
		r, g, b, _ := img.At(x, 0).RGBA()
		fmt.Printf("%d %d %d @ %d\n", int(r/257), int(g/257), int(b/257), x)
	}
	fmt.Printf("min x,y: %d,%d\n", img.Bounds().Min.X, img.Bounds().Min.Y)
	fmt.Printf("max x,y: %d,%d\n", img.Bounds().Max.X, img.Bounds().Max.Y)

}
