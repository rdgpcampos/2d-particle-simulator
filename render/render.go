package main

import (
	//"fmt"
	"image"
	"os"
	"strconv"

	//"time"

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

	// Execute function below and comment out rest of run() to display a basic animation
	TestRun()

}

func main() {
	pixelgl.Run(run)
}

func TestRun() {
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
	const numParticles = 2 
	var delta_x [numParticles]float64
	var delta_y [numParticles]float64

	t := 0
	positions, err := util.ParseFileToLines(lib.HomePath+"lib/test-simple-run.txt")

	//last := time.Now()
	for !win.Closed() {
		//dt := time.Since(last).Seconds()
		//last = time.Now()

		for i := 0; i < numParticles; i++ {
			id, _ := strconv.ParseInt(util.SplitPositionLine(positions[t])[1], 10, 64)
			delta_x[id], _ = strconv.ParseFloat(util.SplitPositionLine(positions[t])[4],64)
			delta_y[id], _ = strconv.ParseFloat(util.SplitPositionLine(positions[t])[5],64)
			//fmt.Printf(fmt.Sprint(id)+"\n")
		}

		t = (t+9)%len(positions)


		win.Clear(colornames.Black)

		mat := pixel.IM
		mat2 := pixel.IM
		mat = mat.Moved(win.Bounds().Center().Add(pixel.V(delta_x[0]*100,delta_y[0]*100)))
		mat2 = mat2.Moved(win.Bounds().Center().Add(pixel.V(delta_x[1]*100,delta_y[1]*100)))
		mat = mat.Scaled(win.Bounds().Center(),0.05)
		mat2 = mat2.Scaled(win.Bounds().Center(),0.05)
		sprite.Draw(win, mat)
		sprite.Draw(win, mat2)

		win.Update()
	}
}

