package util

import (
	"github.com/rdgpcampos/parallel-2d-particle-simulator/particle"
)

func RemoveParticleByIndex(s []*particle.Particle, index int) []*particle.Particle {
    ret := make([]*particle.Particle, 0)
    ret = append(ret, s[:index]...)
    return append(ret, s[index+1:]...)
}