package particle

import "testing"

func TestCheckParticleType(t *testing.T) {
	result := CheckParticleType("Argon")
	if result != true {
		t.Errorf("Result was incorrect, got false, want true.")
	}

	result = CheckParticleType("Test")
	if result != false {
		t.Errorf("Result was incorrect, got true, want false.")
	}
}