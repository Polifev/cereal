package nav

import "github.com/Polifev/cereal/math"

type Navigation interface {
	Compute(src, dest math.Vector) Path
}
