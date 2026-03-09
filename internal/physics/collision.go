package physics

import (
	"github.com/neerajsurjaye/spen/internal/enums"
	"github.com/neerajsurjaye/spen/internal/model"
	"github.com/neerajsurjaye/spen/internal/smath"
)

type CollisionManifold struct {
	BodyA *model.Body
	BodyB *model.Body

	Normal smath.Vec2
	Penetration float32
	Contact smath.Vec2
}

type CollisionFunc func(a, b model.WorldObject) *CollisionManifold
var collisionTable map[enums.ColliderType]map[enums.ColliderType]CollisionFunc 

func InitCollisionTable(){	
	collisionTable = make(map[enums.ColliderType]map[enums.ColliderType]CollisionFunc)

	collisionTable[enums.ColliderCircle] = make(map[enums.ColliderType]CollisionFunc)
	collisionTable[enums.ColliderWall] = make(map[enums.ColliderType]CollisionFunc)

	collisionTable[enums.ColliderCircle][enums.ColliderCircle] = circleVsCircle
	collisionTable[enums.ColliderCircle][enums.ColliderWall] = circleVsWall
	collisionTable[enums.ColliderWall][enums.ColliderWall] = wallVsWall
}

func circleVsCircle(a, b model.WorldObject) *CollisionManifold{
	return nil
}

func circleVsWall(a, b model.WorldObject) *CollisionManifold{
	return nil
}

func wallVsWall(a, b model.WorldObject) *CollisionManifold{
	return nil
}


func CheckCollision(a, b model.WorldObject) *CollisionManifold{

	colA := a.Type()
	colB := b.Type()


	if fn := collisionTable[colA][colB]; fn != nil{
		return fn(a, b)
	}

	if fn := collisionTable[colB][colA]; fn != nil{
		return fn(b, a)
	}
	return nil
}