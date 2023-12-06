package util

import (
	"testing"
	"github.com/rdgpcampos/parallel-2d-particle-simulator/lib"
)

func TestParseFileToLines(t *testing.T) {
	lines, _ := ParseFileToLines(lib.HomePath+"lib/test-positions.txt")

	if lines[0] != "1 0.0000 2.0000" || lines[5] != "6 5.0000 64.0000" {
		t.Errorf("File was not parsed properly")
	}
}

func TestRemoveParticleByIndex(t *testing.T) {
	var integers []*int
	var (
		a = 1
		b = 2
		c = 3
	)
	integers = append(integers, &a, &b, &c)
	particles := RemoveParticleByIndex[int](integers, 1)

	if *particles[0] != 1 || *particles[1] != 3 {
		t.Errorf("Slice was not modified as expected")
	}
}