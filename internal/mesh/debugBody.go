package mesh

import "github.com/neerajsurjaye/spen/internal/model"

func DebugBody(debugLines *DebugLines, object model.WorldObject, color *model.Color) {
	if object.IsStatic(){
		return
	}

	body := object.GetBody()

	base := object.GetTransform().Position
	velocity := body.Velocity
	
	if velocity.Magnitude() > 100{
		velocity = velocity.Normalize()
		velocity = velocity.Multiply(100)
	}

	end := base.Add(velocity)

	debugLines.AddArrow(base , end, color)
}