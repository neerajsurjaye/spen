package physics

import (
	"github.com/neerajsurjaye/spen/internal/mesh"
	"github.com/neerajsurjaye/spen/internal/model"
)

func DebugCollision(debugLines *mesh.DebugLines, manifolds []*CollisionManifold) {

	for _, m := range manifolds {
		normal := m.Normal
		contactPoint := m.Contact

		end := contactPoint.Add(normal.Multiply(100))

		debugLines.AddArrow(contactPoint, end, model.GetColor(0, 0, 1, 0))
	}

}
