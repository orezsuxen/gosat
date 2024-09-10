package physic

import (
	vec "local/gosat/vector"
)

type collisionInfo struct {
	Distance   float64
	Direction  vec.Vector
	Seperation vec.Vector
	Contained  int // -1=A contained +1=B contained 0=no containmend
	PositionA  vec.Vector
	PositionB  vec.Vector
	SecondTest bool
}

func CheckCollisionCircleCircle(
	a Circle,
	b Circle,
	reverseOrder bool,
) (hasCollision bool, info collisionInfo) {
	retval := collisionInfo{}
	radiusTotal := a.Collider.Radius*a.Transform.Size + b.Collider.Radius*b.Transform.Size
	distanceDelta := a.Transform.Position.Dist(&b.Transform.Position)
	if distanceDelta > radiusTotal {
		return false, retval
	}
	// collision direction
	retval.Direction = vec.Add(&a.Transform.Position, &b.Transform.Position)
	retval.Direction.Unit()

	// distance between colliders
	retval.Distance = distanceDelta
	spaceDelta := distanceDelta - radiusTotal
	sep := vec.MulScalar(&retval.Direction, spaceDelta)
	retval.Seperation = vec.MulScalar(&sep, -1.0)

	// positions
	retval.PositionA = a.Transform.Position
	retval.PositionB = b.Transform.Position

	return true, retval

}
