package mesh

import "github.com/neerajsurjaye/spen/internal/model"

func DebugCircleBody(debugLines *DebugLines, circle *model.Circle, color *model.Color) {
	body := circle.Body

	base := circle.Transform.Position
	velocity := body.Velocity
	
	if velocity.Magnitude() > 100{
		velocity = velocity.Normalize()
		velocity = velocity.Multiply(100)
	}

	end := base.Add(velocity)

	debugLines.AddArrow(base , end, color)
}