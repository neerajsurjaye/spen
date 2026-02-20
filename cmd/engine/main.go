package main

import (
	"fmt"

	"github.com/neerajsurjaye/spen/internal/builder"
	"github.com/neerajsurjaye/spen/internal/model"
	"github.com/neerajsurjaye/spen/internal/simulation"
	"github.com/neerajsurjaye/spen/internal/state"
)

func main() {

	state.InitEngine()
	state.InitWorld()
	engine := state.EngineInstance()
	world := state.WorldInstance()

	world.SetBackground(0.8, 1, 0.7, 1)

	engine.UiInfo = model.UiInfo{Width: 800, Height: 800, Title: "Spectres Physics Engine"}
	engine.GlfwState.Window = simulation.InitGlfw()
	simulation.InitOpenGl()

	circleBuilder := builder.GetCircleBuilder().
										Compute().
										SetShaders("shader/circle/circle.frag", "shader/circle/circle.vert")

	circle1 := circleBuilder.Build()
	circle2 := circleBuilder.Build()
	circle3 := circleBuilder.Build()


	fmt.Println(circle1 , " " , circle2)

	circle2.SetColor(1, 0.5, 1, 1)
	circle2.SetTransform(0.5, 0.2, 0.3)

	circle1.SetColor(0.5, 1, 1, 1)
	circle1.SetTransform(-0.5, -0.8, 0.6)

	circle3.SetColor(0.8, 0.7, 0.9, 1)
	circle3.SetTransform(0.1, -0.3, 0.4)

	world.AddCircle(circle1)
	world.AddCircle(circle2)
	world.AddCircle(circle3)

	simulation.StartLoop()
	
}