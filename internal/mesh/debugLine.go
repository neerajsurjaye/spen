package mesh

import (
	"github.com/go-gl/gl/all-core/gl"
	"github.com/neerajsurjaye/spen/internal/model"
	"github.com/neerajsurjaye/spen/internal/renderer"
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

	gl.VertexAttribPointer(0, 2, gl.FLOAT, false, 6 * 4, nil)
	gl.VertexAttribPointer(1, 4, gl.FLOAT, false, 6 * 4, gl.PtrOffset(2 * 4))

	gl.EnableVertexAttribArray(0)
	gl.EnableVertexAttribArray(1)
	//Unbinds array buffer
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)
}

func (d *DebugLines) AddLine(
	x1, y1 float32,
	x2, y2 float32,
	color *model.Color,
){
	d.vertices = append(d.vertices, x1, y1, color.R,color.G, color.B, color.A)
	d.vertices = append(d.vertices, x2, y2, color.R,color.G, color.B, color.A)

}

func (d *DebugLines) Draw(rc *renderer.RenderContext){
	if !rc.DebugDraw{
		return
	}

	if len(d.vertices) == 0 {
		return
	}

	gl.UseProgram(d.shader)

	mvp := rc.OrthoProjection.	
				Mul4(rc.ViewMat)
	
	mvpLoc := gl.GetUniformLocation(d.shader, gl.Str("uMVP\x00"))
	gl.UniformMatrix4fv(mvpLoc, 1, false, &mvp[0])

	gl.BindVertexArray(d.vao)

	gl.BindBuffer(gl.ARRAY_BUFFER, d.vbo)

	gl.BufferData(gl.ARRAY_BUFFER, len(d.vertices) * 4, gl.Ptr(d.vertices), gl.DYNAMIC_DRAW)

	gl.DrawArrays(gl.LINES, 0 , int32(len(d.vertices) / 6))

	gl.BindVertexArray(0)

	//Lines will be added every frame
	d.vertices = d.vertices[:0]
}