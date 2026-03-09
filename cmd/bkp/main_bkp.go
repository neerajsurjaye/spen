// package main

// import (
// 	"fmt"
// 	"log"
// 	"runtime"
// 	"strings"

// 	"github.com/go-gl/gl/all-core/gl"
// 	"github.com/go-gl/glfw/v3.3/glfw"
// )

// const(
// 	width = 800
// 	height = 800
// )

// var triangle []float32 = []float32{
// 	0, 0.5 ,0,
// 	-0.5, -0.5, 0,
// 	0.5, -0.5, 0,
// }

// var vertexShaderSource = `
// 	#version 410
// 	in vec3 vp;
// 	void main(){
// 		gl_Position = vec4(vp, 1.0);
// 	}
// ` + "\x00"

// var fragmentShaderSource = `
// 	#version 410
// 	out vec4 frag_colour;
// 	void main(){
// 		frag_colour = vec4(1,1,1,1);
// 	}
// ` + "\x00"


// func init(){
// 	runtime.LockOSThread()
// }

// func main() {
// 	fmt.Println("Starting SPEN")

// 	window := initGlfw()
// 	defer glfw.Terminate()

// 	prog := initOglProgram()

// 	vao := makeVao(triangle)


// 	for !window.ShouldClose(){

// 		draw(vao, window, prog)
// 	}

// }

// func initGlfw() *glfw.Window{
// 	if err := glfw.Init(); err != nil{
// 		panic(err)
// 	}
	
// 	glfw.WindowHint(glfw.Resizable, glfw.False)
//     glfw.WindowHint(glfw.ContextVersionMajor, 4) 
//     glfw.WindowHint(glfw.ContextVersionMinor, 1)
//     glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
//     glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

// 	window, err := glfw.CreateWindow(width, height, "Spectres Physics Engine", nil, nil)

// 	if err != nil{
// 		panic(err)
// 	}

// 	window.MakeContextCurrent()

// 	return window
// }

// func initOglProgram() uint32{

// 	if err := gl.Init(); err != nil{
// 		log.Fatalln(err)
// 	}

// 	//Just Logs
// 	version := gl.GoStr(gl.GetString(gl.VERSION))	
// 	renderer := gl.GoStr(gl.GetString(gl.RENDERER))
// 	vendor := gl.GoStr(gl.GetString(gl.VENDOR))

// 	log.Println("OpenGL Version :", version)
// 	log.Println("Renderer       :", renderer)
// 	log.Println("Vendor         :", vendor)

// 	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
// 	if err != nil{
// 		panic(err)
// 	}

// 	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
// 	if err != nil{
// 		panic(err)
// 	}

// 	prog := gl.CreateProgram()
// 	gl.AttachShader(prog, vertexShader)
// 	gl.AttachShader(prog, fragmentShader)
// 	gl.LinkProgram(prog)
// 	return prog
// }


// func draw(vao uint32, window *glfw.Window, program uint32){
// 	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
// 	gl.UseProgram(program)

// 	gl.BindVertexArray(vao)
// 	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(triangle)/ 3))

// 	window.SwapBuffers()
// 	glfw.PollEvents()
// }



// func makeVao(points []float32) uint32{
// 	var vbo uint32 //Vertex buffer object

// 	// VBO tells take this much of memory and copy this data to gpu
// 	//Creates a GPU buffer
// 	gl.GenBuffers(1, &vbo)
// 	//Any operation on ArrayBuffer refers to this VBO.
// 	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
// 	gl.BufferData( 					//Push Data to GPU
// 		gl.ARRAY_BUFFER,			//Upload as array buffer 
// 		4*len(points),				//Allocate this much mem in gpu. 4 is size of float in bytes. So 4 * number of elements.
// 		gl.Ptr(points),				//Copies to GPU by taking its pointer
// 		gl.STATIC_DRAW)				//Tells data wont change very much

// 	var vao uint32
// 	gl.GenVertexArrays(1, &vao)									//Creates Vertex Array object
// 	gl.BindVertexArray(vao)										//Makes it the active VAO
// 	gl.EnableVertexAttribArray(0)								//Attribute loc as 0 **Research more on it**
// 	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)							//Binds VBO as as currnt array buffer
// 	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)		//0 attributes location. Components per vertex(x,y,z). Then datatype, Normalize false correct for floats,Stride 0 = open gl caculates itself , offset nil start at beginning of buffer.

// 	return vao
// }


// /**
// 	Compiles shader
// */
// func compileShader(source string, shaderType uint32) (uint32, error){
// 	shader := gl.CreateShader(shaderType)

// 	csources, free := gl.Strs(source)
// 	gl.ShaderSource(shader, 1, csources, nil)
// 	free()
// 	gl.CompileShader(shader)

// 	var status int32
// 	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)

// 	if status == gl.FALSE{
// 		var logLength int32
// 		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

// 		log := strings.Repeat("\x00", int(logLength + 1))
// 		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

// 		return 0, fmt.Errorf("Failed to compil %v: %v", source, log)
// 	}

// 	return shader, nil
// }

