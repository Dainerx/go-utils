package main

import (
	"bufio"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
	"time"

	"github.com/boombuler/barcode"
	Ean "github.com/boombuler/barcode/ean"
	"github.com/disintegration/imaging"
	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
)

const (
	fontFolder = "/home/dainer/go/src/github.com/Dainerx/go-utils/pkg/barcode/resource/font/"
	// constants from https://internationalbarcodes.com/ean-13-specifications/
	widthImage  = 200
	heightImage = 120

	widthEancode  = 90  // px
	heightEanCode = 12  // px - 11.640944882
	widthBarcode  = 119 // px - 118.48818898
	heightBarcode = 86  // px - 86.362204724 - excluding the number at the bottom

	fileExt = ".png"
)

var numbersRunes = []rune("1234567890")

// This is completly wrong due to check sum.
func InitEans(fileName string) ([]string, error) {
	eanSeed, err := os.OpenFile(fileName, os.O_RDONLY, 0666)
	if err != nil {
		return nil, err
	}

	eans := make([]string, 0)
	scanner := bufio.NewScanner(eanSeed)
	for scanner.Scan() {
		line := scanner.Text()
		eans = append(eans, line)
	}

	return eans, nil
}

// Encode Ean in one image
func EncodeEan(ean string) (image.Image, error) {
	draw2d.SetFontFolder(fontFolder)
	// Create an empty white image
	codeImg := image.NewRGBA(image.Rect(0, 0, widthEancode, heightEanCode))
	white := color.RGBA{255, 255, 255, 255}
	draw.Draw(codeImg, codeImg.Bounds(), &image.Uniform{white}, image.ZP, draw.Src)
	// Write EAN at the bottom
	gc := draw2dimg.NewGraphicContext(codeImg)
	gc.FillStroke()
	gc.SetFontData(draw2d.FontData{Name: "BarcodeFont", Family: draw2d.FontFamilyMono, Style: draw2d.FontStyleBold | draw2d.FontStyleItalic})
	gc.SetFillColor(image.Black)
	gc.SetFontSize(10)
	gc.FillStringAt(ean, 0, 10)

	// Create a barcode
	barCode, err := Ean.Encode(ean)
	if err != nil {
		return nil, err
	}
	// Scale it
	barCodeScaled, err := barcode.Scale(barCode, widthBarcode, heightBarcode)
	if err != nil {
		return nil, err
	}

	// Final image
	finalImage := imaging.New(widthImage, heightImage, color.NRGBA{255, 255, 255, 255})

	// Paste the barcode to the final image
	finalImage = imaging.Paste(finalImage, barCodeScaled, image.Pt(50, 10))

	// Paste the number to the final image
	finalImage = imaging.Paste(finalImage, codeImg, image.Pt(64, 100))

	return finalImage, nil
}

func WriteImage(img image.Image, fileName string) error {
	// encode the barcode as png
	file, err := os.Create(fileName + fileExt)
	if err != nil {
		return err
	}

	defer file.Close()

	return png.Encode(file, img)
}

type BarcodeImage struct {
	ean string
	img image.Image
}

func main() {

	eans, err := InitEans("seed-clean.txt")
	if err != nil {
		log.Fatalf("Failed to read eans from seed: %v", err)

	}
	imageChannel := make(chan *BarcodeImage, len(eans))

	t := time.Now()
	for _, ean := range eans {
		go func(ean string) {
			barcodeImage, err := EncodeEan(ean)
			if err != nil {
				imageChannel <- &BarcodeImage{ean, nil}
			}

			imageChannel <- &BarcodeImage{ean, barcodeImage}
		}(ean)
	}

	count := 0
	encoded := 0
	for barcodeimg := range imageChannel {
		if barcodeimg.img == nil {
			log.Printf("Failed to encode %s\n", barcodeimg.ean)
		} else {
			WriteImage(barcodeimg.img, barcodeimg.ean)
			log.Printf("Successfully encoded %s\n", barcodeimg.ean)
			encoded++
		}

		count++
		if count == len(eans) {
			close(imageChannel)
		}
	}

	log.Printf("Finished encoded %d, time taken: %v", encoded, time.Since(t))
}
