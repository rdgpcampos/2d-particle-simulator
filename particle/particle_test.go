package particle

import (
	"fmt"
	"testing"

	"github.com/rdgpcampos/parallel-2d-particle-simulator/lib"
	"github.com/rdgpcampos/parallel-2d-particle-simulator/util"
)

func TestCheckParticleType(t *testing.T) {
	result, _ := CheckParticleType("Argon")
	if result != true {
		t.Errorf("Result was incorrect, got false, want true.")
	}

	result, _ = CheckParticleType("Test")
	if result != false {
		t.Errorf("Result was incorrect, got true, want false.")
	}
}

func TestLogSystem(t *testing.T) {
	p1, _ := New("Argon", 1.0, 0.0, 0.0, 0.0, 0.0)
	p2, _ := New("Argon", 1.0, 1.0, 0.0, 0.0, 0.0)

	var particles []*Particle

	particles = append(particles, p1, p2)

	filepath := lib.HomePath+"/lib/test-output-log.txt"

	LogSystem(3.0, particles, filepath)

	lines, _ := util.ParseFileToLines(filepath)
	if lines[len(lines)-1] != "3.00 1 Argon 1.00 1 0" {
		fmt.Printf(lines[len(lines)-1]+"\n")
		t.Errorf("Failed to print line")
	}
}