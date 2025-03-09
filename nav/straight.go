package nav

import "github.com/Polifev/cereal/math"

type Straight struct{}

func (*Straight) Compute(src, dest math.Vector) Path {
	return Path{
		nodes: []math.Vector{dest},
	}
}
