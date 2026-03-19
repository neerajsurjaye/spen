package simulation

import (
	"time"

	"github.com/go-gl/gl/all-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/neerajsurjaye/spen/internal/force"
	"github.com/neerajsurjaye/spen/internal/model"
	"github.com/neerajsurjaye/spen/internal/physics"
	"github.com/neerajsurjaye/spen/internal/smath"
	"github.com/neerajsurjaye/spen/internal/state"
	"github.com/neerajsurjaye/spen/internal/utils"
)

func StartLoop() {

	engine := state.EngineInstance()
	world := state.WorldInstance()
	window := engine.GlfwState.Window

	running := true
	lastTime := time.Now()
	accumulator := 0.0

	w := state.WorldInstance()
	// log := logger.NewLogger(100)

	var targetFrameTime time.Duration
	if TARGET_FPS > 0 {
		targetFrameTime = time.Second / time.Duration(TARGET_FPS)
	}

	for running && !window.ShouldClose() {

		frameStart := time.Now()
		now := time.Now()
		frameTime := now.Sub(lastTime).Seconds()
		lastTime = now

		if frameTime >= MAX_FRAME_TIME {
			/**
			Say I pause the program. Then the next frametime would be very large.
			If program paused for 10s. Frametime would be 10s.
			So this clapms the frametime to some max value.
			*/
			frameTime = MAX_FRAME_TIME
		}

		if !engine.IsPaused || engine.StepSim {
			accumulator += frameTime

			if engine.StepSim {
				accumulator = FIXED_DT
			}

			for accumulator >= FIXED_DT {
				resetCollisoin(w.WorldObjects)
				performPhysics(w.WorldObjects, FIXED_DT)
				// checkCollision(w.WorldObjects)
				manifolds := detectCollision(w.WorldObjects)
				world.CollisionManifold = manifolds
				resolveCollision(manifolds)

				accumulator -= FIXED_DT
			}

			engine.StepSim = false
		}

		running = draw()

		if TARGET_FPS > 0 {
			frameDuration := time.Since(frameStart)

			/**
			If frame came before 1/120 second. Then sleep for the remaining time.
			*/
			if frameDuration < targetFrameTime {
				time.Sleep(targetFrameTime - frameDuration)
			}
		}

	}

}

func draw() bool {

	world := state.WorldInstance()
	engine := state.EngineInstance()
	window := state.EngineInstance().GlfwState.Window

	window.SwapBuffers()
	glfw.PollEvents()

	if window.GetKey(glfw.KeyW) == glfw.Press {
		world.Camera.MoveDelta(0, 10)
	}

	if window.GetKey(glfw.KeyS) == glfw.Press {
		world.Camera.MoveDelta(0, -10)
	}

	if window.GetKey(glfw.KeyA) == glfw.Press {
		world.Camera.MoveDelta(-10, 0)
	}

	if window.GetKey(glfw.KeyD) == glfw.Press {
		world.Camera.MoveDelta(+10, 0)
	}

	if window.GetKey(glfw.KeyR) == glfw.Press {
		world.Camera.X = 0
		world.Camera.Y = 0
	}

	if utils.IsKeyPressed(window, glfw.KeyI) {
		world.DebugDraw = !world.DebugDraw
	}

	if utils.IsKeyPressed(window, glfw.KeyEscape) {
		return false
	}

	if utils.IsKeyPressed(window, glfw.KeyP) {
		engine.IsPaused = !engine.IsPaused
	}

	if utils.IsKeyPressed(window, glfw.KeyL) {
		if engine.IsPaused {
			engine.StepSim = true
		}
	}

	Render()

	return true
}

func Render() {

	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	world := state.WorldInstance()
	world.Draw()

}

func performPhysics(worldObjects []model.WorldObject, dt float32) {

	for i := range worldObjects {
		object := worldObjects[i]

		if object.IsStatic() {
			continue
		}

		body := object.GetBody()
		gravity := force.NewGravity()

		body.AddForce(gravity.GetForce(body.Mass))
		body.IntegrateVelocity(dt)

		//Integrate Position
		newPosition := object.GetTransform().Position.Add(body.Velocity.Multiply(dt))
		object.SetPosition(&newPosition)
	}
}

/*Depriciated*/
func checkCollision(worldObjects []model.WorldObject) {

	for i := 0; i < len(worldObjects); i++ {
		for j := i + 1; j < len(worldObjects); j++ {

			if worldObjects[i].GetAABB().IsColliding(worldObjects[j].GetAABB()) {

				worldObjects[i].GetAABB().Colliding = true
				worldObjects[j].GetAABB().Colliding = true

				if !worldObjects[i].IsStatic() {
					worldObjects[i].GetBody().SetVelocity(smath.Vec2{X: 0, Y: 0})
				}
				if !worldObjects[j].IsStatic() {
					worldObjects[j].GetBody().SetVelocity(smath.Vec2{X: 0, Y: 0})
				}
			}
		}
	}
}

func detectCollision(worldObjects []model.WorldObject) []*physics.CollisionManifold {

	var manifolds []*physics.CollisionManifold

	for i := 0; i < len(worldObjects); i++ {
		for j := i + 1; j < len(worldObjects); j++ {
			if worldObjects[i].GetAABB().IsColliding(worldObjects[j].GetAABB()) {
				m := physics.CheckCollision(worldObjects[i], worldObjects[j])

				if m != nil {
					manifolds = append(manifolds, m)
				}
			}
		}
	}

	return manifolds
}

func resolveCollision(manifolds []*physics.CollisionManifold) {
	for _, m := range manifolds {
		physics.ResolveImpulse(m.BodyA, m.BodyB, *m)
	}

	for _, m := range manifolds {
		physics.PositionCorrection(m.BodyA, m.BodyB, *m)
	}
}

func resetCollisoin(worldObjects []model.WorldObject) {
	for i := range worldObjects {
		worldObjects[i].GetAABB().Colliding = false
	}
}
