package force

import (
	"github.com/parallel-2d-particle-simulator/particle"
)

type Force struct {
	X float64
	Y float64
}

type ForceCalc func(*particle.Particle,[]*particle.Particle) (*Force, error)

func forceSum(forces ...*Force) *Force {
	totalForce := &Force{}

	if len(forces) == 0 {
		return &Force{}
	}

	for _, f := range forces {
		totalForce.X += f.X
		totalForce.Y += f.Y
	}

	return totalForce
}



