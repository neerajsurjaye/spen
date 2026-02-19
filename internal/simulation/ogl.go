package simulation

import (
	"log"

	"github.com/go-gl/gl/all-core/gl"
	"github.com/neerajsurjaye/spen/internal/state"
)

func InitOpenGl() {
	if err := gl.Init(); err != nil {
		log.Fatalln(err)
	}

	uiInfo := state.EngineInstance().UiInfo

	version := gl.GoStr(gl.GetString(gl.VERSION))	
	renderer := gl.GoStr(gl.GetString(gl.RENDERER))
	vendor := gl.GoStr(gl.GetString(gl.VENDOR))

	log.Println("OpenGL Version :", version)
	log.Println("Renderer       :", renderer)
	log.Println("Vendor         :", vendor)

	//TODo: use global state here
	gl.Viewport(0,0 , int32(uiInfo.Width), int32(uiInfo.Height))
	gl.ClearColor(0.1, 0.2, 0.8, 0)
}

func Render() {

	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	engine := state.EngineInstance()

	engine.Circle.Draw()

}