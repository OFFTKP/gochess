package main

import (
	"fmt"
	"image"
	"os"

	"github.com/OFFTKP/gochess/core"
	"github.com/fogleman/gg"
)

var squareSize float64 = 64
var squareSizeHalf float64 = squareSize / 2
var imagesMap [12]*image.Image
var imagesLoaded bool = false

func getImageFromFilePath(filePath string) image.Image {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	image, _, err := image.Decode(f)
	return image
}

func outputCurrentImage(board *core.Board) {
	loadImages()
	dc := gg.NewContext(int(squareSize)*8, int(squareSize)*8)
	var x, y float64
	var black bool = true
	for y = 7; y >= 0; y-- {
		for x = 0; x < 8; x++ {
			dc.DrawRectangle(x*squareSize, y*squareSize, squareSize, squareSize)
			if black {
				dc.SetRGB(0, 0.3, 0)
			} else {
				dc.SetRGB(0, 0.6, 0)
			}
			dc.Fill()
			curIndex := int(x) + (int(7-y) * 8)
			var curPiece uint8 = 0xFF
			for i := 0; i < 12; i++ {
				for j := uint8(0); j < 12; j++ {
					curBB := board.PieceBBmap[j]
					if ((*curBB) & (1 << curIndex)) != 0 {
						curPiece = j
						break
					}
				}
			}
			if curPiece != 0xFF {
				dc.DrawImageAnchored(*imagesMap[curPiece], int(x*squareSize+squareSizeHalf), int(y*squareSize+squareSizeHalf), 0.5, 0.5)
			}
			black = !black
		}
		black = !black
	}
	dc.SavePNG("out.png")
}

func loadImages() {
	if !imagesLoaded {
		wpawnImage := getImageFromFilePath("data/wpawn.png")
		wknightImage := getImageFromFilePath("data/wknight.png")
		wbishopImage := getImageFromFilePath("data/wbishop.png")
		wrookImage := getImageFromFilePath("data/wrook.png")
		wqueenImage := getImageFromFilePath("data/wqueen.png")
		wkingImage := getImageFromFilePath("data/wking.png")
		bpawnImage := getImageFromFilePath("data/bpawn.png")
		bknightImage := getImageFromFilePath("data/bknight.png")
		bbishopImage := getImageFromFilePath("data/bbishop.png")
		brookImage := getImageFromFilePath("data/brook.png")
		bqueenImage := getImageFromFilePath("data/bqueen.png")
		bkingImage := getImageFromFilePath("data/bking.png")
		imagesMap = [12]*image.Image{
			&wpawnImage, &wknightImage, &wbishopImage, &wrookImage, &wqueenImage, &wkingImage,
			&bpawnImage, &bknightImage, &bbishopImage, &brookImage, &bqueenImage, &bkingImage,
		}
		imagesLoaded = true
	}
}

func main() {
	var board core.Board
	board.LoadFen("rnbqkbnr/pppppppp/8/8/4P3/8/1PPPPPPP/RNBQKBNR w KQkq - 0 1")
	ret := board.SinglePushTargets(0)
	ret |= board.DoublePushTargets(0)
	fmt.Printf("%016x", ret)
	outputCurrentImage(&board)
}
