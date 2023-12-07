package force

import (
	"github.com/rdgpcampos/parallel-2d-particle-simulator/particle"
	"github.com/rdgpcampos/parallel-2d-particle-simulator/lib"
	"errors"
	"math"
)

func GravitationalForceAtParticle(curParticle *particle.Particle, particles []*particle.Particle) (*Force, error) {

	err := error(nil)
	totalForce := &Force{}
	curForce := &Force{}
	var r float64

	for _, p := range particles {
		if curParticle.Pos_x == p.Pos_x && curParticle.Pos_y == p.Pos_y {
			err = errors.New("Force is undefined at these coordinates")
			return &Force{0,0}, err
		}

		// distance from position x,y to current particle
		r = math.Sqrt(math.Pow(curParticle.Pos_x-p.Pos_x,2)+math.Pow(curParticle.Pos_y-p.Pos_y,2))

		// force acting on current particle exerted by other p
		curForce = &Force{lib.GravitationalConstant*p.Mass*curParticle.Mass*(p.Pos_x - curParticle.Pos_x)/(math.Pow(r,3)),
							lib.GravitationalConstant*p.Mass*curParticle.Mass*(p.Pos_y - curParticle.Pos_y)/(math.Pow(r,3))}


		totalForce = forceSum(totalForce, curForce)
	}

	return totalForce, err
}