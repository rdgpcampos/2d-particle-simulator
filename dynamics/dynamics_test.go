package dynamics

import (
	"testing"
	"reflect"
)

import (
	"github.com/parallel-2d-particle-simulator/particle"
)

func TestMove(t *testing.T) {
	particle := particle.New("Argon", 1.0, 0.0, 0.0, 1.0, 2.0, 0.0, 0.0)

	Move(particle, 1.0)

	if reflect.ValueOf(particle).Elem().FieldByName("Pos_x").Float() != 1.0 {
		t.Errorf("X position was not calculated properly")
	}

	if reflect.ValueOf(particle).Elem().FieldByName("Pos_y").Float() != 2.0 {
		t.Errorf("Y position was not calculated properly")
	}
	
}




