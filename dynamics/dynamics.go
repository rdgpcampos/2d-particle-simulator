package dynamics

import (
	"github.com/parallel-2d-particle-simulator/particle"
	"reflect"
)

func Move(particle *particle.Particle, delta_t float64) {
	particle_elem := reflect.ValueOf(particle).Elem()

	// update position of particle
	particle_elem.FieldByName("Pos_x").
		SetFloat(particle_elem.FieldByName("Pos_x").Float()+delta_t*particle_elem.FieldByName("Vel_x").Float())
	particle_elem.FieldByName("Pos_y").
		SetFloat(particle_elem.FieldByName("Pos_y").Float()+delta_t*particle_elem.FieldByName("Vel_y").Float())
}




