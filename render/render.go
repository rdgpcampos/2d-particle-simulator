package main

import (
	"image"
	"os"

	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"github.com/rdgpcampos/parallel-2d-particle-simulator/lib"
)

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	pic, err := loadPicture(lib.ImagePath+"circle2.png")
	if err != nil {
		panic(err)
	}

	sprite := pixel.NewSprite(pic, pic.Bounds())

	win.Clear(colornames.Black)
	delta_x := 0.0
	delta_y := 0.0

	for !win.Closed() {
		delta_x += 2
		delta_y += 4

		win.Clear(colornames.Black)

		mat := pixel.IM
		mat = mat.Moved(win.Bounds().Center().Add(pixel.V(delta_x,delta_y)))
		mat = mat.Scaled(win.Bounds().Center(),0.05)
		sprite.Draw(win, mat)

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}

