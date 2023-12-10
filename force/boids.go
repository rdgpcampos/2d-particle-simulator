package force

import (
	"math"

	"github.com/rdgpcampos/parallel-2d-particle-simulator/lib"
	"github.com/rdgpcampos/parallel-2d-particle-simulator/particle"
)

func BoidForceAtParticle(curParticle *particle.Particle, particles []*particle.Particle) (*Force, error) {
	var err error
	totalForce := &Force{}

	totalForce = forceSum(totalForce, separate(curParticle,particles))
	totalForce = forceSum(totalForce, cohesion(curParticle,particles))
	totalForce = forceSum(totalForce, align(curParticle,particles))

	return totalForce, err
}

func separate(curParticle *particle.Particle, particles []*particle.Particle) *Force {
	totalForce := &Force{}
	curForce := &Force{}
	var r float64

	for _, p := range particles {
		if curParticle.Pos_x == p.Pos_x && curParticle.Pos_y == p.Pos_y {
			return &Force{0,0}
		}

		// distance from position x,y to current particle
		r = math.Sqrt(math.Pow(curParticle.Pos_x-p.Pos_x,2)+math.Pow(curParticle.Pos_y-p.Pos_y,2))
		if r > lib.BoidSeparateDistance {
			continue
		}

		// force acting on current particle exerted by other p
		curForce = &Force{lib.BoidSeparateConstant*(p.Vel_x - curParticle.Vel_x)/(r/lib.BoidSeparateDistance),
							lib.BoidSeparateConstant*(p.Vel_y - curParticle.Vel_y)/(r/lib.BoidSeparateDistance)}

		curForce = forceSum(curForce,&Force{lib.BoidSeparateConstant*(curParticle.Pos_x - p.Pos_x)/(r/lib.BoidSeparateDistance),
							lib.BoidSeparateConstant*(curParticle.Pos_y - p.Pos_y)/(r/lib.BoidSeparateDistance)})


		totalForce = forceSum(totalForce, curForce)
	}


	return totalForce
}

func cohesion(curParticle *particle.Particle, particles []*particle.Particle) *Force {
	totalForce := &Force{}
	curForce := &Force{}
	var r float64
	n := len(particles)

	for _, p := range particles {
		// distance from position x,y to current particle
		r = math.Sqrt(math.Pow(curParticle.Pos_x-p.Pos_x,2)+math.Pow(curParticle.Pos_y-p.Pos_y,2))
		if r > lib.BoidSize {
			continue
		}

		// force acting on current particle exerted by other p
		curForce = &Force{lib.BoidCohesionConstant*(p.Pos_x - curParticle.Pos_x)/float64(n),lib.BoidCohesionConstant*(p.Pos_y - curParticle.Pos_y)/float64(n)}

		totalForce = forceSum(totalForce, curForce)
	}

	return totalForce
}

func align(curParticle *particle.Particle, particles []*particle.Particle) *Force {
	totalForce := &Force{}
	curForce := &Force{}
	var r float64
	n := len(particles)

	for _, p := range particles {
		// distance from position x,y to current particle
		r = math.Sqrt(math.Pow(curParticle.Pos_x-p.Pos_x,2)+math.Pow(curParticle.Pos_y-p.Pos_y,2))
		if r > lib.BoidSize {
			continue
		}

		// force acting on current particle exerted by other p
		curForce = &Force{lib.BoidAlignConstant*p.Vel_x/float64(n),lib.BoidAlignConstant*p.Vel_y/float64(n)}

		totalForce = forceSum(totalForce, curForce)
	}

	return totalForce
}