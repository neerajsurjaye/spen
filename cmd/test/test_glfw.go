package main

import (
	"log"
	"runtime"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func init(){
	// Locks the thread to one OS thread. Open gl issue.
	runtime.LockOSThread()
}

func main() {
	//Loads plaform specific windowing code
	err := glfw.Init()
	

	if(err != nil){
		panic(err)
	}

	defer glfw.Terminate()

	window, err := glfw.CreateWindow(800, 800, "Glfw Test", nil, nil)

	if err != nil{

		panic(err)
	}

	window.MakeContextCurrent()

	if err := gl.Init(); err != nil{
		log.Fatalln(err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	renderer := gl.GoStr(gl.GetString(gl.RENDERER))
	vendor := gl.GoStr(gl.GetString(gl.VENDOR))

	log.Println("OpenGL Version :", version)
	log.Println("Renderer       :", renderer)
	log.Println("Vendor         :", vendor)

	for !window.ShouldClose(){
		window.SwapBuffers()
		glfw.PollEvents()
	}
}