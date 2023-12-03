package particle

import (
	"bufio"
	"os"
	"log"
	"slices"
	"github.com/parallel-2d-particle-simulator/lib"
)

type Particle struct {
	Pos_x float64
	Pos_y float64
	Vel_x float64
	Vel_y float64
	Acc_x float64
	Acc_y float64
	Mass float64
	Ptype string
}


func New(Ptype string, 
			Mass float64, 
			Pos_x float64, 
			Pos_y float64, 
			Vel_x float64, 
			Vel_y float64, 
			Acc_x float64, 
			Acc_y float64) *Particle {
	
	if Mass <= 0 {
		panic("New: particle mass cannot be a negative number")
	}

	if !CheckParticleType(Ptype) {
		panic("New: particle type does not exist")
	}

	return &Particle{Ptype: Ptype, Mass: Mass, Pos_x: Pos_x, Pos_y: Pos_y, Vel_x: Vel_x, Vel_y: Vel_y, Acc_x: Acc_x, Acc_y: Acc_y}
}


func CheckParticleType(ptype string) bool {

	// open file describing particle types
	file, err := os.Open(lib.HomePath+"/lib/particle-types.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// read particle types into string array
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	defer func() {
		if err = scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}()

	// check if particle type exists
	return slices.Contains(lines, ptype)
}