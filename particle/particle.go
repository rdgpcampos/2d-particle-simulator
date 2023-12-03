package particle

import (
	"bufio"
	"os"
	"log"
	"slices"
)

type particle struct {
	pos_x float32
	pos_y float32
	vel_x float32
	vel_y float32
	mass float32
	ptype string
}


func New(ptype string, mass float32, pos_x float32, pos_y float32) *particle {
	
	if mass <= 0 {
		panic("New: particle mass cannot be a negative number")
	}

	if !CheckParticleType(ptype) {
		panic("New: particle type does not exist")
	}

	return &particle{ptype: ptype, mass: mass, pos_x: pos_x, pos_y: pos_y}
}


func CheckParticleType(ptype string) bool {

	// open file describing particle types
	file, err := os.Open("particle-types.txt")
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

// next write simpletic particle position change