package physics

import (
	"fmt"
	"math"

	"github.com/neerajsurjaye/spen/internal/enums"
	"github.com/neerajsurjaye/spen/internal/model"
	"github.com/neerajsurjaye/spen/internal/smath"
)

type CollisionManifold struct {
	BodyA model.WorldObject
	BodyB model.WorldObject

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
	circleA, okA := a.(*model.Circle)
	circleB, okB := b.(*model.Circle)

	if !okA || !okB{
		return nil
	}

	posA := circleA.GetTransform().Position
	posB := circleB.GetTransform().Position
	collisionDir := posB.Subtract(posA)
	dist := collisionDir.Magnitude()
	radiusSum := circleA.GetRadius() + circleB.GetRadius()

	if dist >= radiusSum{
		return nil
	}

	normal := collisionDir.Normalize()

	manifold := CollisionManifold{
		BodyA: a,
		BodyB: b,	
	}

	if dist != 0{
		manifold.Penetration = radiusSum - dist
		manifold.Normal = normal
		manifold.Contact = posA.Add(normal.Multiply(circleA.GetRadius()))
	}else{
		manifold.Penetration = circleA.GetRadius()
		manifold.Normal = smath.Vec2{X: 1, Y: 0}
		manifold.Contact = a.GetTransform().Position
	}
	
	return &manifold
}

func circleVsWall(a, b model.WorldObject) *CollisionManifold{

	typeA := a.GetColliderType()
	typeB := b.GetColliderType()

	if !((typeA == enums.ColliderCircle && typeB == enums.ColliderWall) || (typeB == enums.ColliderCircle && typeA == enums.ColliderWall)){
		return nil
	}

	var circle *model.Circle
	var wall *model.Wall 
	var okA, okB bool
	if(typeA == enums.ColliderCircle){
		circle, okA= a.(*model.Circle)
		wall, okB = b.(*model.Wall)
	}else{
		circle, okA = b.(*model.Circle)
		wall, okB = a.(*model.Wall)
	}

	if !okA || !okB{
		return nil
	}

	circlePos := circle.Transform.Position
	rectPos := wall.Transform.Position
	rot := wall.Transform.Radians

	hx := wall.GetTransform().Scale.X
	hy := wall.GetTransform().Scale.Y

	//Converting circe coords to rectangle coords
	local := circlePos.Subtract(rectPos)

	cos  := float32(math.Cos(float64(-rot)))
	sin  := float32(math.Sin(float64(-rot)))

	localX := local.X*cos - local.Y * sin
	localY := local.X*sin + local.Y * cos 

	localPoint := smath.NewVec2(localX, localY)

	fmt.Println(localPoint)
	// closestX := 

	return nil

}

func wallVsWall(a, b model.WorldObject) *CollisionManifold{
	return nil
}


func CheckCollision(a, b model.WorldObject) *CollisionManifold{

	colA := a.GetColliderType()
	colB := b.GetColliderType()


	if fn := collisionTable[colA][colB]; fn != nil{
		return fn(a, b)
	}

	if fn := collisionTable[colB][colA]; fn != nil{
		return fn(b, a)
	}
	return nil
}

func CheckAABBColl(a, b model.WorldObject) *CollisionManifold{

	p := a.GetAABB()
	q := b.GetAABB()

	collision := p.MaxX >= q.MinX &&
		p.MinX <= q.MaxX &&
		p.MaxY >= q.MinY &&
		p.MinY <= q.MaxY

	if !collision{
		return nil
	}

	overlapX := min(p.MaxX, q.MaxX) - max(p.MinX, q.MinX)
	overlapY := min(p.MaxY, q.MaxY) - max(p.MinY, q.MinY) 

	centerAX := (p.MinX + p.MaxX) * 0.5
	centerBX := (q.MinX + q.MaxX) * 0.5
	centerAY := (p.MinY + p.MaxY) * 0.5
	centerBY := (q.MinY + q.MaxY) * 0.5

	manifold := &CollisionManifold{
		BodyA: a,
		BodyB: b,
	}

	if overlapX <  overlapY{
		if centerAX < centerBX{
			manifold.Normal = smath.NewVec2(-1, 0)
		}else{
			manifold.Normal = smath.NewVec2(1, 0)
		}
		manifold.Penetration = overlapX
	}else{
		if centerAY < centerBY{
			manifold.Normal = smath.NewVec2(0, -1)
		}else{
			manifold.Normal = smath.NewVec2(0 , 1)
		}
		manifold.Penetration = overlapY
	}

	return manifold

}