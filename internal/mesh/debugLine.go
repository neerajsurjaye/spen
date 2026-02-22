package mesh

import (
	"github.com/go-gl/gl/all-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type DebugLines struct {
	vao uint32
	vbo uint32

	vertices []float32

	shader uint32
}

func (d *DebugLines) Init(shader uint32) {
	d.shader = shader

	gl.GenVertexArrays(1, &d.vao)
	gl.GenBuffers(1, &d.vbo)

	gl.BindVertexArray(d.vao)

	gl.BindBuffer(gl.ARRAY_BUFFER, d.vbo)

	gl.VertexAttribPointer(0, 2, gl.FLOAT, false, 2 * 4, nil)

	gl.EnableVertexAttribArray(0)
	//Unbinds array buffer
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)
}

func (d *DebugLines) AddLine(
	x1, y1 float32,
	x2, y2 float32,
){
	d.vertices = append(d.vertices, x1, y1, x2, y2)
}

func (d *DebugLines) Draw(projection mgl32.Mat4, view mgl32.Mat4){
	if len(d.vertices) == 0 {
		return
	}

	gl.UseProgram(d.shader)

	mvp := projection.	
				Mul4(view)
	
	mvpLoc := gl.GetUniformLocation(d.shader, gl.Str("uMVP\x00"))
	gl.UniformMatrix4fv(mvpLoc, 1, false, &mvp[0])

	gl.BindVertexArray(d.vao)

	gl.BindBuffer(gl.ARRAY_BUFFER, d.vbo)

	gl.BufferData(gl.ARRAY_BUFFER, len(d.vertices) * 4, gl.Ptr(d.vertices), gl.DYNAMIC_DRAW)

	gl.DrawArrays(gl.LINES, 0 , int32(len(d.vertices) / 2))

	gl.BindVertexArray(0)

	//Lines will be added every frame
	d.vertices = d.vertices[:0]
}