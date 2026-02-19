package main

import (
	"github.com/neerajsurjaye/spen/internal/mesh"
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

	var circle mesh.CircleInfo = mesh.GetCircle(0,0,1)
	circle.ComputeCircleVAO()
	circle.AttachShader("shader/circle/circle.frag", "shader/circle/circle.vert")
	circle.SetColor(1, 1, 1, 1)

	engine.Circle = circle

	simulation.StartLoop()
	
}