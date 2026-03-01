package main

import (
	"math"
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

// ====================================================================

	circleBuilder := builder.GetCircleBuilder().Compute()
	circleMesh := circleBuilder.BuildObject()

	wallBuilder := builder.GetWallBuilder().Compute()
	wallMesh := wallBuilder.BuildObject()

	greatColor := model.Color{R: 1, G: 1, B: 1, A: 1}
	circle1 := model.GetCircle(0, 0 , 100, 0, greatColor)
	circle2 := model.GetCircle(500.2, 400 , 230, math.Pi / 2, greatColor)
	circle3 := model.GetCircle(-250, 400 , 120, math.Pi / 3, greatColor)

	circle1.SetBody(0, 100, 10)
	circle2.SetBody(0, 0, 100)
	circle3.SetBody(0, 0, 100)

	world.AddObject(circle1, circleMesh)
	world.AddObject(circle2, circleMesh)
	world.AddObject(circle3, circleMesh)

	groundColor := model.GetColor(0.3, 1, 0.3, 1)
	ground := model.GetWall(0, -100, 100, 30, math.Pi / 2, *groundColor)
	world.AddObject(ground, wallMesh)

	movingSquare := model.GetWall(30, -30, 40, 40, math.Pi / 3, *groundColor)
	movingSquare.SetBody(10, 50, 100)
	world.AddObject(movingSquare, wallMesh)

	simulation.StartLoop()
	
}