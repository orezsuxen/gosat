package physic

import (
	vec "local/gosat/vector"
)

type PhysicWorld struct {
	dimension vec.Vector
	minForce  float64
	data      []PhysBody
	gravity   vec.Vector
}

func (w *PhysicWorld) updatePosition(delta float64) {

	for _, p := range w.data {
		if p.GetBodyDynamic() {
			t := p.GetTransform()
			b := p.GetBody()
			//apply gravity
			deltaGravity := w.gravity
			deltaGravity.MulScalar(delta)
			b.Force.Add(&deltaGravity)
			// apply force
			deltaForce := b.Force
			deltaForce.MulScalar(delta)
			t.Position.Add(&deltaForce)
			//? apply drag

			// limit
			if b.Force.Magn() < w.minForce {
				b.Force = vec.Vector{X: 0.0, Y: 0.0}
			}
		}
	}
}

func (w *PhysicWorld) killOutside() {
	for i, p := range w.data {
		t := p.GetTransform()
		if t.Position.X > w.dimension.X ||
			t.Position.X < -w.dimension.X ||
			t.Position.Y > w.dimension.Y ||
			t.Position.Y < -w.dimension.Y {
			w.data[i] = w.data[len(w.data)-1]
			w.data = w.data[:len(w.data)-1]

		}
	}
}

//TODO: collision checking
