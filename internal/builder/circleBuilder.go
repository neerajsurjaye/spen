package builder

import (
	"github.com/go-gl/gl/all-core/gl"
	"github.com/neerajsurjaye/spen/internal/mesh"
	"github.com/neerajsurjaye/spen/internal/utils"
)

type CircleBuilder struct {
	// circle mesh.CircleInfo
	fragShaderPath string
	vertexShaderPath string

	circleVert []float32

	// Calculate these things here
	vao uint32
	vbo uint32
	prog uint32
}


var circleVert []float32 = []float32{
							-1, -1,
							1, -1,
							1, 1,
							-1, 1,
						}

func GetCircleBuilder() *CircleBuilder{
	return &CircleBuilder{
		circleVert: circleVert,
	}
}

func (cb *CircleBuilder) SetShaders(fragShaderPath string, vertexShaderPath string) *CircleBuilder{
	cb.fragShaderPath = fragShaderPath
	cb.vertexShaderPath = vertexShaderPath
	cb.attachShader(cb.fragShaderPath, cb.vertexShaderPath)
	return cb
}

func (cb *CircleBuilder) Compute() *CircleBuilder{
	cb.computeCircleVAO()
	cb.SetShaders("shader/circle/circle.frag", "shader/circle/circle.vert")
	return cb
}	

func (cb *CircleBuilder) Build() mesh.CircleInfo{
	if cb.vao == 0 || cb.vbo == 0{
		panic("Circle VAO not implemented. Call CircleBuilder.Compute() before CircleBuilder.Build()")
	}
	if cb.prog == 0{
		panic("Shader not implemented. Call CircleBuilder.SetShaders() before CircleBuilder.Build()")
	}

	return mesh.GetCircle(cb.vao, cb.vbo, cb.prog)
}

func (cb *CircleBuilder) computeCircleVAO(){
	var vao, vbo uint32

	gl.GenVertexArrays(1, &vao)
	gl.GenBuffers(1, &vbo)

	gl.BindVertexArray(vao)

	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(
		gl.ARRAY_BUFFER,
		len(cb.circleVert) * 4,
		gl.Ptr(cb.circleVert),
		gl.STATIC_DRAW,
	)

	gl.VertexAttribPointer(0, 2, gl.FLOAT, false, 0, nil)
	gl.EnableVertexAttribArray(0)

	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)
	cb.vao = vao
	cb.vbo = vbo
}

func (cb *CircleBuilder) attachShader(fragShaderPath string, vertexShaderPath string){
	//TODO: Check this part
	cb.prog = utils.CreateProgram(fragShaderPath, vertexShaderPath)
}

