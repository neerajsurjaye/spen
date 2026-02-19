package simulation

import (
	"time"

	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/neerajsurjaye/spen/internal/state"
)

func StartLoop() {

	engine := state.EngineInstance()
	window := engine.GlfwState.Window

	running := true
	lastTime := time.Now()
	accumulator := 0.0

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

		accumulator += frameTime	


		// log.Log("Render loop at : " , now , accumulator)
		for accumulator >= FIXED_DT{
			//Perform physics
			// log.Log("Perfoming Physics at : " , now , accumulator)
			accumulator -= FIXED_DT
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

	window := state.EngineInstance().GlfwState.Window
	window.SwapBuffers()
	glfw.PollEvents()

	Render()

	if window.GetKey(glfw.KeyEscape) == glfw.Press{
		return false
	}
	return true
}