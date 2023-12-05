package force

type Force struct {
	X float64
	Y float64
}

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

