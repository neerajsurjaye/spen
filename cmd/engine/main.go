package main

import (
	"math"
	"runtime"

	"math/rand"

	"github.com/neerajsurjaye/spen/internal/builder"
	"github.com/neerajsurjaye/spen/internal/model"
	"github.com/neerajsurjaye/spen/internal/physics"
	"github.com/neerajsurjaye/spen/internal/simulation"
	"github.com/neerajsurjaye/spen/internal/state"
	"github.com/neerajsurjaye/spen/internal/utils"
)

func init() {
	// Locks the thread to one OS thread. Open gl issue.
	runtime.LockOSThread()
}

func main() {

	state.InitEngine()
	state.InitWorld()
	state.InitInputState()
	physics.InitCollisionTable()

	engine := state.EngineInstance()
	world := state.WorldInstance()

	engine.UiInfo = model.UiInfo{Width: 1600, Height: 800, Title: "Spectres Physics Engine"}
	engine.GlfwState.Window = simulation.InitGlfw()

	world.SetCamera(0, 0)
	world.SetBackground(0.2, 0.2, 0.2, 1)

	simulation.InitOpenGl()

	world.DebugLines.Init(utils.CreateProgram("shader/line/line.frag", "shader/line/line.vert"))
	world.Grid.Init(100, 1000)

	// ====================================================================

	circleBuilder := builder.GetCircleBuilder().Compute()
	circleMesh := circleBuilder.BuildObject()

	wallBuilder := builder.GetWallBuilder().Compute()
	wallMesh := wallBuilder.BuildObject()

	greatColor := model.Color{R: 1, G: 1, B: 1, A: 1}
	circle1 := model.GetCircle(0, 300, 50, 0, greatColor)
	circle2 := model.GetCircle(0.001, 0, 50, math.Pi/2, greatColor)
	circle3 := model.GetCircle(-100, 300, 10, 0, greatColor)
	circle4 := model.GetCircle(-80, 0, 10, math.Pi/2, greatColor)
	circle5 := model.GetCircle(150, 300, 10, 0, greatColor)
	circle6 := model.GetCircle(100, 0, 10, math.Pi/2, greatColor)
	// circle3 := model.GetCircle(450, 400, 100, math.Pi/3, greatColor)

	circle1.SetBody(0, 0, 1000, 1)
	circle2.SetBody(20, 200, 10, 0.8)
	circle3.SetBody(0, 100, 1000, 1)
	circle4.SetBody(60, 20, 100, 1)
	circle5.SetBody(40, 30, 10, 1)
	circle6.SetBody(0, 1000, 1000, 1)

	for i := range 500 {
		// color := float32(i / 10)

		size := rand.Float32() * 10

		circlex := model.GetCircle(0, 0, size, 0, *model.GetColor(rand.Float32(), rand.Float32(), rand.Float32(), 1))
		circlex.SetBody(0, 0, float32(50*i), 0.1)
		world.AddObject(circlex, circleMesh)
	}

	// circle3.SetBody(0, -200, 1000, 0.8)

	world.AddObject(circle1, circleMesh)
	world.AddObject(circle2, circleMesh)
	world.AddObject(circle3, circleMesh)
	world.AddObject(circle4, circleMesh)
	world.AddObject(circle5, circleMesh)
	world.AddObject(circle6, circleMesh)
	// world.AddObject(circle3, circleMesh)

	groundColor := model.GetColor(0.3, 1, 0.3, 1)
	ground := model.GetWall(200, -200, 800, 20, math.Pi/100, *groundColor)
	world.AddObject(ground, wallMesh)

	ground = model.GetWall(-200, -400, 800, 20, -math.Pi/100, *groundColor)
	world.AddObject(ground, wallMesh)

	ground = model.GetWall(200, -600, 800, 20, math.Pi/100, *groundColor)
	world.AddObject(ground, wallMesh)

	ceil := model.GetWall(0, 400, 800, 20, 0, *groundColor)
	world.AddObject(ceil, wallMesh)

	wallLeft := model.GetWall(-900, 0, 20, 500, 0, *groundColor)
	world.AddObject(wallLeft, wallMesh)

	wallRight := model.GetWall(900, 0, 20, 500, 0, *groundColor)
	world.AddObject(wallRight, wallMesh)

	// movingSquare := model.GetWall(300, -100, 20, 40, math.Pi/3, *groundColor)
	// movingSquare.SetBody(0, 0, 100, 0.1)
	// world.AddObject(movingSquare, wallMesh)

	simulation.StartLoop()

}
