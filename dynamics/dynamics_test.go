package dynamics

import (
	"testing"
	"github.com/parallel-2d-particle-simulator/particle"
)

func TestMove(t *testing.T) {
	particle, _ := particle.New("Argon", 1.0, 0.0, 0.0, 1.0, 2.0)

	Move(particle, 1.0)

	if particle.Pos_x != 1.0 {
		t.Errorf("X position was not calculated properly")
	}

	if particle.Pos_y != 2.0 {
		t.Errorf("Y position was not calculated properly")
	}
}




