package physic

import (
	"math"

	vec "local/gosat/vector"
)

// REM:Parts
type CircleCollider struct {
	Radius float64
}

type PolygonCollider struct {
	verts []vec.Vector
}

type TransformData struct {
	Position vec.Vector
	Rotation float64
	Size     float64
}

type BodyData struct {
	dynamic bool
	Mass    float64
	Force   vec.Vector
	Drag    vec.Vector
}

// REM: interface
type PhysBody interface {
	GetTransform() *TransformData
	GetBody() *BodyData
	GetBodyDynamic() bool
}

// REM: Circle
type Circle struct {
	Transform TransformData
	Collider  CircleCollider
	Body      BodyData
}

// interface func
func (c *Circle) GetTransform() *TransformData {
	return &c.Transform
}

// interface func
func (c *Circle) GetBody() *BodyData {
	return &c.Body
}

// interface func
func (c *Circle) GetBodyDynamic() bool {
	return c.Body.dynamic
}

// REM: Triangle
type Triangle struct {
	Transform TransformData
	Collider  PolygonCollider
	Body      BodyData
}

// interface func
func (t *Triangle) GetTransform() *TransformData {
	return &t.Transform
}

// interface func
func (t *Triangle) GetBody() *BodyData {
	return &t.Body
}

// interface func
func (t *Triangle) GetBodyDynamic() bool {
	return t.Body.dynamic
}

//REM: functions

func buildPolygonCollider(verts int) PolygonCollider {
	angleDelta := 2.0 * math.Pi / float64(verts)
	coll := make([]vec.Vector, verts)
	for i := range verts {
		coll[i] = vec.VectorFromRot(float64(i) * angleDelta)
	}

	return PolygonCollider{
		verts: coll,
	}
}

func buildCircleCollider() CircleCollider {

	return CircleCollider{
		1.0,
	}
}
