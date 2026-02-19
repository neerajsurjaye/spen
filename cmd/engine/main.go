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
	engine := state.EngineInstance()

	engine.UiInfo = model.UiInfo{Width: 800, Height: 800, Title: "Spectres Physics Engine"}
	engine.GlfwState.Window = simulation.InitGlfw()
	simulation.InitOpenGl()


	circleBuilder := builder.GetCircleBuilder().
										Compute().
										SetShaders("shader/circle/circle.frag", "shader/circle/circle.vert")

	circle1 := circleBuilder.Build()
	circle2 := circleBuilder.Build()


	fmt.Println(circle1 , " " , circle2)

	circle2.SetColor(1, 0.5, 1, 1)
	engine.Circle = circle2

	simulation.StartLoop()
	
}