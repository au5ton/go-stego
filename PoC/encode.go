package main

/*

Proof of Concept
================
Encoding hidden data into PNG data

*/

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"strconv"
)

type Image interface {
	image.Image
	Set(x, y int, c color.Color)
}

var input = flag.String("i", "", "Input file name of PNG")
var output = flag.String("o", "", "Output file name of PNG")

func main() {
	flag.Parse()

	if *input == "" {
		//log.Fatal("input cannot be empty")
	}
	if *output == "" {
		//log.Fatal("output cannot be empty")
	}

	var str = "00000"
	var runes = []byte(str)
	for i := 0; i < len(runes); i++ {
		fmt.Print(strconv.FormatInt(int64(runes[i]), 2), " ")
	}
	fmt.Println(runes)

}
