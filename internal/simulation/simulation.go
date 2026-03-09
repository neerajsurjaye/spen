package simulation

import (
	"fmt"
	"time"

	"github.com/go-gl/gl/all-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/neerajsurjaye/spen/internal/force"
	"github.com/neerajsurjaye/spen/internal/model"
	"github.com/neerajsurjaye/spen/internal/smath"
	"github.com/neerajsurjaye/spen/internal/state"
)

func StartLoop() {

	engine := state.EngineInstance()
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

	for running && !window.ShouldClose(){

		frameStart := time.Now()
		now := time.Now()
		frameTime := now.Sub(lastTime).Seconds()
		lastTime = now	

		if frameTime >= MAX_FRAME_TIME{
			/**
				Say I pause the program. Then the next frametime would be very large.
				If program paused for 10s. Frametime would be 10s.
				So this clapms the frametime to some max value.
			*/
			frameTime = MAX_FRAME_TIME
		}


		if !engine.IsPaused || engine.StepSim{ 
			accumulator += frameTime
			
			if(engine.StepSim){
				accumulator = FIXED_DT
			}

			for accumulator >= FIXED_DT{
				//Perform physics
				for idx := range(w.WorldObjects){
					performPhysics(w.WorldObjects[idx], FIXED_DT)
				}

				for idx := range(w.WorldObjects){
					checkCollision(w.WorldObjects[idx], w.WorldObjects)
				}
				// log.Log("Perfoming Physics at : " , now , accumulator)
				accumulator -= FIXED_DT
			}

			engine.StepSim = false
		}
		
		running = draw()

		if(TARGET_FPS > 0){
			frameDuration := time.Since(frameStart)

			/**
				If frame came before 1/120 second. Then sleep for the remaining time.
			*/
			if frameDuration < targetFrameTime{
				time.Sleep(targetFrameTime - frameDuration)
			}
		}

	}

}

func draw() bool{

	world := state.WorldInstance()
	engine := state.EngineInstance()
	window := state.EngineInstance().GlfwState.Window

	window.SwapBuffers()
	glfw.PollEvents()


	if window.GetKey(glfw.KeyW) == glfw.Press{
		world.Camera.MoveDelta(0, 10)
	}

	if window.GetKey(glfw.KeyS) == glfw.Press{
		world.Camera.MoveDelta(0, -10)
	}

	if window.GetKey(glfw.KeyA) == glfw.Press{
		world.Camera.MoveDelta(-10, 0)
	}

	if window.GetKey(glfw.KeyD) == glfw.Press{
		world.Camera.MoveDelta(+10, 0)
	}

	if window.GetKey(glfw.KeyR) == glfw.Press{
		world.Camera.X = 0
		world.Camera.Y = 0
	}

	if IsKeyPressed(window, glfw.KeyI){
		world.DebugDraw = !world.DebugDraw
	}

	if IsKeyPressed(window, glfw.KeyEscape){
		return false
	}

	if IsKeyPressed(window, glfw.KeyP){
		engine.IsPaused = !engine.IsPaused
	}

	if IsKeyPressed(window, glfw.KeyL){
		if engine.IsPaused{
			engine.StepSim = true
		}
	}

	Render()

	return true
}


func Render() {

	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	// engine := state.EngineInstance()
	world := state.WorldInstance()
	world.Draw()

}

func performPhysics(objects model.WorldObject, dt float32){
	if objects.IsStatic(){
		return
	}

	body := objects.GetBody()

	gravity := force.NewGravity()

	body.AddForce(gravity.GetForce(body.Mass))

	body.Integrate(dt)

	newPosition := objects.GetTransform().Position.Add(body.Velocity.Multiply(dt))
	fmt.Println("Diff Position ", objects.GetTransform().Position.Subtract(newPosition))
	objects.SetPosition(&newPosition)
}

func checkCollision(object model.WorldObject, worldObjects []model.WorldObject){

	for i := range(worldObjects){
		if(object != worldObjects[i] && object.GetAABB().IsColliding(worldObjects[i].GetAABB()) && !object.IsStatic()){
			object.GetBody().SetVelocity(smath.Vec2{X : 0, Y: 0})
		}	
	}
}

func IsKeyPressed(window *glfw.Window, key glfw.Key) bool{
	inputState := state.InputStateInstance()

	currentAction := window.GetKey(key)
	prevAction := inputState.PrevAction[key]

	inputState.PrevAction[key] = currentAction

	return currentAction == glfw.Press && prevAction == glfw.Release
}