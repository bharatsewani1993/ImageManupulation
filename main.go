package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"log"
	"os"

	"github.com/disintegration/imaging"
)

func main() {
	CreateImages()
	fmt.Println("Image Created")
}

//CreateImages merge png image on white background.
func CreateImages() {

	//Decode background and pic
	backGround, err := os.Open("background.png")
	productpic, err := os.Open("4.png")
	backGroundd, err := png.Decode(backGround)
	productpicd, err := png.Decode(productpic)

	if err != nil {
		fmt.Println(err)
	}

	//Draw productpic on white backgroud
	offset := image.Pt(300, 200)
	b := backGroundd.Bounds()
	canvas := image.NewRGBA(b)
	draw.Draw(canvas, b, backGroundd, image.ZP, draw.Src)
	draw.Draw(canvas, productpicd.Bounds().Add(offset), productpicd, image.ZP, draw.Over)
	out, err := os.Create("./done/done4.jpg")
	if err != nil {
		fmt.Println(err)
	}
	var opt jpeg.Options
	opt.Quality = 100
	jpeg.Encode(out, canvas, &opt)

	//Dwrawn Image
	src, err := imaging.Open("./done/done4.jpg")
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	//Adjust Brightness of Image
	dst := imaging.AdjustBrightness(src, 20)

	//Save Image
	err = imaging.Save(dst, "./done/bright4.jpg")
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}

}
