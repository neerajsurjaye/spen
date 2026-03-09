package enums

type ColliderType int

const (
	ColliderCircle ColliderType = iota
	ColliderWall   ColliderType = iota
	ColliderAABB   ColliderType = iota
)