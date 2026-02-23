package main

import (
	"fmt"
	"runtime"

	"github.com/neerajsurjaye/spen/internal/builder"
	"github.com/neerajsurjaye/spen/internal/model"
	"github.com/neerajsurjaye/spen/internal/simulation"
	"github.com/neerajsurjaye/spen/internal/state"
	"github.com/neerajsurjaye/spen/internal/utils"
)

func init(){
	// Locks the thread to one OS thread. Open gl issue.
	runtime.LockOSThread()
}

func main() {

	state.InitEngine()
	state.InitWorld()
	engine := state.EngineInstance()
	world := state.WorldInstance()

	engine.UiInfo = model.UiInfo{Width: 800, Height: 800, Title: "Spectres Physics Engine"}
	engine.GlfwState.Window = simulation.InitGlfw()

	world.SetCamera(0, 0)
	world.SetBackground(0.2, 0.2, 0.2, 1)

	simulation.InitOpenGl()

	world.DebugLines.Init(utils.CreateProgram("shader/line/line.frag" , "shader/line/line.vert"))
	world.Grid.Init(100, 500) 


	circleBuilder := builder.GetCircleBuilder().Compute()

	circleMesh := circleBuilder.Build()

	fmt.Println(circleMesh)

	world.CircleMesh = circleMesh

	greatColor := model.Color{R: 0.5, G: 1, B: 1, A: 1}
	circle1 := model.GetCircle(0, 0 , 100, greatColor)
	circle2 := model.GetCircle(500.2, 400 , 230, greatColor)
	circle3 := model.GetCircle(-250, 400 , 120, greatColor)
	world.AddCircle(circle1)
	world.AddCircle(circle2)
	world.AddCircle(circle3)

	simulation.StartLoop()
	
}