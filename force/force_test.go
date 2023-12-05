package force

import (
	"fmt"
	"testing"
	"github.com/rdgpcampos/parallel-2d-particle-simulator/particle"
)

func TestForceSum(t *testing.T) {
	f1 := &Force{1,2}
	f2 := &Force{0,3}

	f3 := forceSum(f1,f2)
	f4 := &Force{1,5}

	if f3.X != f4.X || f3.Y != f4.Y {
		t.Errorf("Force was not calculated property")
	}
}

func TestForceAtParticle(t *testing.T) {
	p1, _ := particle.New("Argon", 1.0, 0.0, 0.0, 0.0, 0.0)
	p2, _ := particle.New("Argon", 1.0, 3.0, 4.0, 0.0, 0.0)
	p3, _ := particle.New("Argon", 1.0, 5.0, 12.0, 0.0, 0.0)

	var particles []*particle.Particle

	particles = append(particles, p2, p3)

	f1, _ := GravitationalForceAtParticle(p1, particles)

	if f1.X != 3.0/125 + 5.0/13/13/13 || f1.Y != 4.0/125 + 12.0/13/13/13 {
		t.Errorf(fmt.Sprintf("%f :force was not calculated properly", f1.Y),)
	}

}