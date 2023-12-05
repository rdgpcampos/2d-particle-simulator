package dynamics

import (
	"github.com/parallel-2d-particle-simulator/particle"
)

func Move(particle *particle.Particle, delta_t float64) {

	// update position of particle
	particle.Pos_x += delta_t*particle.Vel_x
	particle.Pos_y += delta_t*particle.Vel_y

}




