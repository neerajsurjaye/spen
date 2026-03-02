package model

type AABB struct {
	MinX float32
	MaxX float32
	MinY float32
	MaxY float32

	Colliding bool
}

func (curr *AABB) IsColliding(collider *AABB) bool {

	if curr.MaxX >= collider.MinX &&
		curr.MinX <= collider.MaxX &&
		curr.MaxY >= collider.MinY &&
		curr.MinY <= collider.MaxY {
		curr.Colliding = true
	} else {
		curr.Colliding = false
	}
	return curr.Colliding
}