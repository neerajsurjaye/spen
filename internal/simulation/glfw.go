package simulation

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/neerajsurjaye/spen/internal/state"
)

func InitGlfw() *glfw.Window{

	if err := glfw.Init(); err != nil{
		panic(err)
	}

	uiInfo := state.EngineInstance().UiInfo

	glfw.WindowHint(glfw.Resizable, glfw.False)
    glfw.WindowHint(glfw.ContextVersionMajor, 4) 
    glfw.WindowHint(glfw.ContextVersionMinor, 1)
    glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
    glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(uiInfo.Width, uiInfo.Height, uiInfo.Title, nil, nil)

	if err != nil{
		panic(err)
	}

	window.MakeContextCurrent()

	return window
}