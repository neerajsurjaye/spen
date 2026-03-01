package builder

import (
	"github.com/go-gl/gl/all-core/gl"
	"github.com/neerajsurjaye/spen/internal/mesh"
	"github.com/neerajsurjaye/spen/internal/utils"
)

type WallBuilder struct {
	fragShaderPath string
	vertexShaderPath string

	wallVert []float32

	// Calculate these things here
	vao uint32
	vbo uint32
	prog uint32
}


var wallVert []float32 = []float32{
							-1, -1,
							1, -1,
							1, 1,
							-1, 1,
						}

func GetWallBuilder() *WallBuilder{
	return &WallBuilder{
		wallVert: wallVert,
	}
}

func (cb *WallBuilder) SetShaders(fragShaderPath string, vertexShaderPath string) *WallBuilder{
	cb.fragShaderPath = fragShaderPath
	cb.vertexShaderPath = vertexShaderPath
	cb.attachShader(cb.fragShaderPath, cb.vertexShaderPath)
	return cb
}

func (cb *WallBuilder) Compute() *WallBuilder{
	cb.computeWallVAO()
	cb.SetShaders("shader/wall/wall.frag", "shader/wall/wall.vert")
	return cb
}	

func (cb *WallBuilder) Build() mesh.WallMesh{
	if cb.vao == 0 || cb.vbo == 0{
		panic("Wall VAO not implemented. Call WallBuilder.Compute() before WallBuilder.Build()")
	}
	if cb.prog == 0{
		panic("Shader not implemented. Call WallBuilder.SetShaders() before WallBuilder.Build()")
	}

	return mesh.GetWall(cb.vao, cb.vbo, cb.prog)
}

func (cb *WallBuilder) computeWallVAO(){
	var vao, vbo uint32

	gl.GenVertexArrays(1, &vao)
	gl.GenBuffers(1, &vbo)

	gl.BindVertexArray(vao)

	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(
		gl.ARRAY_BUFFER,
		len(cb.wallVert) * 4,
		gl.Ptr(cb.wallVert),
		gl.STATIC_DRAW,
	)

	gl.VertexAttribPointer(0, 2, gl.FLOAT, false, 0, nil)
	gl.EnableVertexAttribArray(0)

	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)
	cb.vao = vao
	cb.vbo = vbo
}

func (cb *WallBuilder) attachShader(fragShaderPath string, vertexShaderPath string){
	//TODO: Check this part
	cb.prog = utils.CreateProgram(fragShaderPath, vertexShaderPath)
}

