package dynamics

import (
	"github.com/rdgpcampos/parallel-2d-particle-simulator/particle"
	"github.com/rdgpcampos/parallel-2d-particle-simulator/force"
	"github.com/rdgpcampos/parallel-2d-particle-simulator/util"
)

func Move(particles []*particle.Particle, delta_t float64, forceCalc force.ForceCalc) error {

	var err error
	var forceOnParticle *force.Force

	// update positions of particles
	for _,particle := range particles {
		particle.Pos_x += delta_t*particle.Vel_x + delta_t*delta_t*particle.Acc_x*0.5
		particle.Pos_y += delta_t*particle.Vel_y + delta_t*delta_t*particle.Acc_y*0.5
		particle.Vel_x += particle.Acc_x*delta_t*0.5
		particle.Vel_y += particle.Acc_y*delta_t*0.5
	}

	// update acceleration and velocities of particles
	for i,particle := range particles {
		outsideSystem := util.RemoveParticleByIndex(particles,i)
		forceOnParticle, err = forceCalc(particle, outsideSystem)
		particle.Acc_x = forceOnParticle.X/particle.Mass
		particle.Acc_y = forceOnParticle.Y/particle.Mass
		particle.Vel_x += particle.Acc_x*delta_t*0.5
		particle.Vel_y += particle.Acc_y*delta_t*0.5
	}

	return err
}





