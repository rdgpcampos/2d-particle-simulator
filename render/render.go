package main

import (
	"fmt"
	"image"
	"os"
	"strconv"

	//"time"

	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/rdgpcampos/parallel-2d-particle-simulator/lib"
	"github.com/rdgpcampos/parallel-2d-particle-simulator/util"
	"github.com/rdgpcampos/parallel-2d-particle-simulator/setup"
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
	//TestRun()

	// read setup file
	var file_path string
	var log_path string
	var num_particles int

	fmt.Println("Enter run file path:")
	fmt.Scanln(&file_path)

	num_particles, log_path = setup.RunSetupFile(file_path)

	// pixel boilerplate
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

	var delta_x = make([]float64, num_particles)
	var delta_y = make([]float64, num_particles)

	t := 0
	positions, err := util.ParseFileToLines(log_path)

	// run animation
	for !win.Closed() {

		for i := 0; i < num_particles; i++ {
			id, _ := strconv.ParseInt(util.SplitLine(positions[t])[1], 10, 64)
			delta_x[id], _ = strconv.ParseFloat(util.SplitLine(positions[t])[4],64)
			delta_y[id], _ = strconv.ParseFloat(util.SplitLine(positions[t])[5],64)
			t = (t+1)%len(positions)
		}

		win.Clear(colornames.Black)

		var mats = make([]pixel.Matrix,num_particles)

		for i := 0; i < num_particles; i++ {
			mats[i] = pixel.IM
			mats[i] = mats[i].Moved(win.Bounds().Center().Add(pixel.V(delta_x[i]*100,delta_y[i]*100)))
			mats[i] = mats[i].Scaled(win.Bounds().Center(),0.05)
			sprite.Draw(win, mats[i])
		}

		win.Update()
	}

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
	const numParticles = 3 
	var delta_x [numParticles]float64
	var delta_y [numParticles]float64

	t := 0
	positions, err := util.ParseFileToLines(lib.LogsPath+"test-simple-run-from-script.txt")

	for !win.Closed() {

		for i := 0; i < numParticles; i++ {
			id, _ := strconv.ParseInt(util.SplitLine(positions[t])[1], 10, 64)
			delta_x[id], _ = strconv.ParseFloat(util.SplitLine(positions[t])[4],64)
			delta_y[id], _ = strconv.ParseFloat(util.SplitLine(positions[t])[5],64)
			t = (t+1)%len(positions)
		}

		//t = (t+10)%len(positions)


		win.Clear(colornames.Black)

		mat := pixel.IM
		mat = mat.Moved(win.Bounds().Center().Add(pixel.V(delta_x[0]*100,delta_y[0]*100)))
		mat = mat.Scaled(win.Bounds().Center(),0.05)
		sprite.Draw(win, mat)

		mat2 := pixel.IM
		mat2 = mat2.Moved(win.Bounds().Center().Add(pixel.V(delta_x[1]*100,delta_y[1]*100)))
		mat2 = mat2.Scaled(win.Bounds().Center(),0.05)
		sprite.Draw(win, mat2)

		mat3 := pixel.IM
		mat3 = mat3.Moved(win.Bounds().Center().Add(pixel.V(delta_x[2]*100,delta_y[2]*100)))
		mat3 = mat3.Scaled(win.Bounds().Center(),0.05)
		sprite.Draw(win, mat3)

		win.Update()
	}
}

