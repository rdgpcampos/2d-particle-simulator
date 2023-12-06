package main

import (
	"image"
	"os"
	"time"
	"strconv"

	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/rdgpcampos/parallel-2d-particle-simulator/lib"
	"github.com/rdgpcampos/parallel-2d-particle-simulator/util"
	"golang.org/x/image/colornames"

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

	// Execute function below and comment out rest of run() to display a basic animation from test-positions.txt
	TRun()

}

func main() {
	pixelgl.Run(run)
}

func TRun() {
	cfg := pixelgl.WindowConfig{
		Title:  "Test animation",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	win.SetSmooth(true)

	pic, err := loadPicture(lib.ImagePath+"circle2.png")
	if err != nil {
		panic(err)
	}

	sprite := pixel.NewSprite(pic, pic.Bounds())

	win.Clear(colornames.Black)
	delta_x := 0.0
	delta_y := 0.0
	t := 0
	positions, err := util.ParseFileToLines(lib.HomePath+"lib/test-positions.txt")

	last := time.Now()
	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()

		delta_x, _ = strconv.ParseFloat(util.SplitPositionLine(positions[t])[1],64)
		delta_y, _ = strconv.ParseFloat(util.SplitPositionLine(positions[t])[2],64)

		t = (t+1)%len(positions)

		win.Clear(colornames.Black)

		mat := pixel.IM
		mat2 := pixel.IM
		mat = mat.Moved(win.Bounds().Center().Add(pixel.V(delta_x*dt*100,delta_y*dt*100)))
		mat2 = mat2.Moved(win.Bounds().Center().Add(pixel.V(delta_x*dt*1000,delta_y*dt*(-100))))
		mat = mat.Scaled(win.Bounds().Center(),0.05)
		mat2 = mat2.Scaled(win.Bounds().Center(),0.05)
		sprite.Draw(win, mat)
		sprite.Draw(win, mat2)

		win.Update()
	}
}

