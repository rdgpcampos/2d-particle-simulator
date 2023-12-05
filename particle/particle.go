package particle

import (
	"bufio"
	"errors"
	"log"
	"os"
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
		) (*Particle, error) {

	var err error
	var particleTypeExists bool

	if Mass <= 0 {
		err = errors.New("New: particle mass cannot be a negative number")
	}

	// check if particle type is defined before creating particle
	particleTypeExists, err = CheckParticleType(Ptype)
	if !particleTypeExists {
		err = errors.New("New: particle type does not exist")
	}

	return &Particle{Ptype: Ptype, Mass: Mass, Pos_x: Pos_x, Pos_y: Pos_y, Vel_x: Vel_x, Vel_y: Vel_y}, err
}


func CheckParticleType(ptype string) (bool, error) {

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
	return slices.Contains(lines, ptype), err
}