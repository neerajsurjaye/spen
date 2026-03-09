package mesh

import (
	"math"

	"github.com/neerajsurjaye/spen/internal/model"
)

func DebugBody(debugLines *DebugLines, object model.WorldObject, color *model.Color) {
	if object.IsStatic(){
		return
	}

	var mArrowLen float32 = 100
	velocity := object.GetBody().Velocity

	mag := float64(velocity.Magnitude())
	scaledMag := float32(math.Log(mag + 1)) * 10

	if scaledMag > mArrowLen{
		scaledMag = mArrowLen
	}
	scaledVelocity := velocity.Normalize().Multiply(scaledMag)

	base := object.GetTransform().Position
	end := base.Add(scaledVelocity)

	debugLines.AddArrow(base , end, color)
}