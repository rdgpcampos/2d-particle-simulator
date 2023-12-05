package dynamics

import (
	"testing"
	"github.com/parallel-2d-particle-simulator/particle"
	"github.com/parallel-2d-particle-simulator/force"
)

func TestMove(t *testing.T) {
	p1, _ := particle.New("Argon", 1.0, 0.0, 0.0, 0.0, 0.0)
	p2, _ := particle.New("Argon", 1.0, 1.0, 0.0, 0.0, 0.0)

	var particles []*particle.Particle

	particles = append(particles, p1, p2)

	Move(particles, 0.01, force.GravitationalForceAtParticle)
	Move(particles, 0.01, force.GravitationalForceAtParticle)

	if particles[0].Pos_x != 0.000100 {
		t.Errorf("Position was not calculated properly")
	}

	if particles[0].Vel_x < 0.015001 || particles[0].Vel_x > 0.015003 {
		t.Errorf("Velocity was not calculated properly")
	}

	if particles[0].Acc_x < 1.0003 || particles[0].Vel_x > 1.0005 {
		t.Errorf("Acceleration was not calculated properly")
	}
}




