package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"runtime"

	"github.com/neerajsurjaye/spen/internal/builder"
	"github.com/neerajsurjaye/spen/internal/force"
	"github.com/neerajsurjaye/spen/internal/model"
	"github.com/neerajsurjaye/spen/internal/physics"
	"github.com/neerajsurjaye/spen/internal/simulation"
	"github.com/neerajsurjaye/spen/internal/smath"
	"github.com/neerajsurjaye/spen/internal/state"
	"github.com/neerajsurjaye/spen/internal/utils"
)

func init() {
	// Locks the thread to one OS thread. Open gl issue.
	runtime.LockOSThread()
}

func main() {

	if len(os.Args) <= 1 {
		panic("Please provide path to world configuration file.")
	}

	filePath := os.Args[1]

	file, err := os.ReadFile(filePath)

	var config model.Config
	err = json.Unmarshal(file, &config)

	if err != nil {
		panic(err)
	}

	fmt.Println(config)
	fmt.Println(config.Circles)

	state.InitEngine()
	state.InitWorld()
	state.InitInputState()
	physics.InitCollisionTable()

	engine := state.EngineInstance()
	world := state.WorldInstance()

	engine.UiInfo = model.UiInfo{Width: config.Ui.Width, Height: config.Ui.Height, Title: config.Ui.Title}
	engine.GlfwState.Window = simulation.InitGlfw()

	world.SetCamera(0, 0, 1)
	world.SetBackground(config.World.WorldBackground[0], config.World.WorldBackground[1], config.World.WorldBackground[2], config.World.WorldBackground[3])

	if config.World.Gravity != nil {
		gravity := *config.World.Gravity
		world.SetGravity(force.NewGravity(&smath.Vec2{X: gravity[0], Y: gravity[1]}))
	} else {
		world.SetGravity(force.NewGravity(nil))
	}

	simulation.InitOpenGl()

	world.DebugLines.Init(utils.CreateProgram("shader/line/line.frag", "shader/line/line.vert"))
	world.Grid.Init(100, float32(math.Max(float64(config.Ui.Height), float64(config.Ui.Width))))

	// // ====================================================================

	circleBuilder := builder.GetCircleBuilder().Compute()
	circleMesh := circleBuilder.BuildObject()

	wallBuilder := builder.GetWallBuilder().Compute()
	wallMesh := wallBuilder.BuildObject()

	circles := config.Circles

	for _, circleConfig := range circles {
		color := model.GetColor(circleConfig.Color[0], circleConfig.Color[1], circleConfig.Color[2], circleConfig.Color[3])

		tranforms := circleConfig.Transform
		circle := model.GetCircle(tranforms.PosX, tranforms.PosY, tranforms.Radius, 0, *color)

		if circleConfig.Body != nil {
			body := circleConfig.Body
			circle.SetBody(body.VelX, body.VelY, body.Mass, body.Restititon)
		}

		world.AddObject(circle, circleMesh)
	}

	for _, wallConfig := range config.Walls {
		color := model.GetColor(wallConfig.Color[0], wallConfig.Color[1], wallConfig.Color[2], wallConfig.Color[3])

		transforms := wallConfig.Transform
		wall := model.GetWall(transforms.PosX, transforms.PosY, transforms.ScaleX, transforms.ScaleY, smath.DegToRad(transforms.Rotation), *color)

		world.AddObject(wall, wallMesh)
	}

	simulation.StartLoop()
}
