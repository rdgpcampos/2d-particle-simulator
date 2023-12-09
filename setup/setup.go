package setup

import (
	"strconv"

	"github.com/rdgpcampos/parallel-2d-particle-simulator/dynamics"
	"github.com/rdgpcampos/parallel-2d-particle-simulator/force"
	"github.com/rdgpcampos/parallel-2d-particle-simulator/lib"
	"github.com/rdgpcampos/parallel-2d-particle-simulator/particle"
	"github.com/rdgpcampos/parallel-2d-particle-simulator/util"
	"os"
)

func RunSetupFile(filePath string) (int, string) {

	var particles []*particle.Particle
	var cur_t float64 = 0.0
	var delta_t float64
	var total_t float64
	//var log_freq int64
	var set_force force.ForceCalc
	var log_file string = lib.LogsPath+"test-run.txt"
	

	lines, err := util.ParseFileToLines(filePath)
	if err != nil {
		panic("Failed to parse file")
	}

	for _,line := range lines {
		lineAsArray := util.SplitLine(line)
		switch lineAsArray[0] {
		case "create-particle":
			mass,err := strconv.ParseFloat(lineAsArray[2],64)
			if err != nil {
				panic("Mass input is incorrect")
			}

			pos_x,err := strconv.ParseFloat(lineAsArray[3],64)
			if err != nil {
				panic("Position x input is incorrect")
			}

			pos_y,err := strconv.ParseFloat(lineAsArray[4],64)
			if err != nil {
				panic("Position y input is incorrect")
			}

			vel_x,err := strconv.ParseFloat(lineAsArray[5],64)
			if err != nil {
				panic("Velocity x input is incorrect")
			}

			vel_y,err := strconv.ParseFloat(lineAsArray[6],64)
			if err != nil {
				panic("Velocity y input is incorrect")
			}

			particles = append(particles, useParticleNew(lineAsArray[1],mass,pos_x,pos_y,vel_x,vel_y))
		case "set-timestep":
			delta_t,err = strconv.ParseFloat(lineAsArray[1],64)
			if err != nil {
				panic("Timestep input is incorrect")
			}
		case "set-force":
			switch lineAsArray[1] {
			case "gravity":
				set_force = force.GravitationalForceAtParticle
			default:
				panic("Force input is incorrect")
			}
		

		case "log-every":
			//log_freq,_ = strconv.ParseInt(lineAsArray[1],10,64)

		case "log-file":
			log_file = lib.LogsPath + lineAsArray[1]
			if _, err := os.Stat(log_file); err == nil {
				panic("File already exists")
			}

		case "run":
			total_t,err = strconv.ParseFloat(lineAsArray[1],64)
			if err != nil {
				panic("Run input is incorrect")
			}

			for cur_t < total_t {
				dynamics.Move(particles, delta_t, set_force)
				cur_t += delta_t
				particle.LogSystem(cur_t, particles, log_file)
			}
		}
	} 

	return len(particles), log_file
}

// function that returns only a pointer to a new particle instance
func useParticleNew (ptype string, 
					mass float64, 
					pos_x float64, 
					pos_y float64,
					vel_x float64,
					vel_y float64) *particle.Particle {
	p,_ := particle.New(ptype,mass,pos_x,pos_y,vel_x,vel_y)

	return p
}
