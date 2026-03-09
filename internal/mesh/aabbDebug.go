package mesh

import (
	"github.com/neerajsurjaye/spen/internal/model"
)


func DebugAABB(debugLines *DebugLines, aabb *model.AABB, aabbDebugColor *model.Color) {

	if aabb.Colliding{
		aabbDebugColor = model.GetColor(1, 0, 0, 0)
	}

	debugLines.AddLine(aabb.MinX, aabb.MinY, aabb.MaxX, aabb.MinY, aabbDebugColor)
	debugLines.AddLine(aabb.MaxX, aabb.MinY, aabb.MaxX, aabb.MaxY, aabbDebugColor)
	debugLines.AddLine(aabb.MaxX, aabb.MaxY, aabb.MinX, aabb.MaxY, aabbDebugColor)
	debugLines.AddLine(aabb.MinX, aabb.MaxY, aabb.MinX, aabb.MinY, aabbDebugColor)
}	