package particle

import (
	"errors"
	"slices"
	"strconv"

	"github.com/rdgpcampos/parallel-2d-particle-simulator/lib"
	"github.com/rdgpcampos/parallel-2d-particle-simulator/util"
)

type Particle struct {
	Id uint64
	Pos_x float64
	Pos_y float64
	Vel_x float64
	Vel_y float64
	Acc_x float64
	Acc_y float64
	Mass float64
	Ptype string
}

var max_id uint64

func New(Ptype string, 
			Mass float64, 
			Pos_x float64, 
			Pos_y float64, 
			Vel_x float64, 
			Vel_y float64, 
		) (*Particle, error) {

	var err error
	var particleTypeExists bool

	Id := max_id
	max_id += 1

	if Mass <= 0 {
		err = errors.New("New: particle mass cannot be a negative number")
	}

	// check if particle type is defined before creating particle
	particleTypeExists, err = CheckParticleType(Ptype)
	if !particleTypeExists {
		err = errors.New("New: particle type does not exist")
	}

	return &Particle{Id: Id, Ptype: Ptype, Mass: Mass, Pos_x: Pos_x, Pos_y: Pos_y, Vel_x: Vel_x, Vel_y: Vel_y}, err
}

func CheckParticleType(ptype string) (bool, error) {
	// read file containing all valid particle types
	lines, err := util.ParseFileToLines(lib.HomePath+"lib/particle-types.txt")
	// check if particle type exists
	return slices.Contains(lines, ptype), err
}

// log particle info for renderization later
func LogSystem(currTime float64, particles []*Particle, logPath string) error {
	var err error
	
	// Log data: Time ID Type Mass X Y 
	for _,particle := range particles {
		var line string
		line += strconv.FormatFloat(currTime, 'f', 2, 64) + " " + 
				strconv.FormatUint(particle.Id,10) + " " + 
				particle.Ptype + " " + 
				strconv.FormatFloat(particle.Mass, 'f', 2, 64) + " " + 
				strconv.FormatFloat(particle.Pos_x, 'f', -1, 64) + " " +
				strconv.FormatFloat(particle.Pos_y, 'f', -1, 64)
		err = util.AppendLineToFile(line, logPath)
		if err != nil {
			panic(err)
		}
	}

	return err
}