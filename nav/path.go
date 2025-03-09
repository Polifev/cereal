package nav

import "github.com/Polifev/cereal/math"

type Path struct {
	current int
	nodes   []math.Vector
}

func NewPath(nodes []math.Vector) Path {
	return Path{nodes: nodes}
}

func (path *Path) Update(followerPos math.Vector) {
	if path.IsDone() {
		// Should not happen
		return
	}

	if path.Current().EqRough(followerPos) {
		path.Next()
	}
}

func (path *Path) Current() math.Vector {
	return path.nodes[path.current]
}

func (path *Path) Next() bool {
	path.current++
	if path.IsDone() {
		return false
	}
	return true
}

func (path *Path) IsDone() bool {
	return path.current >= len(path.nodes)
}
