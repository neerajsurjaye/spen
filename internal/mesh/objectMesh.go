package mesh

import (
	"github.com/go-gl/gl/all-core/gl"
	"github.com/neerajsurjaye/spen/internal/model"
	"github.com/neerajsurjaye/spen/internal/renderer"
)

type ObjectMesh struct {
	vao uint32
	vbo uint32
	prog uint32
}


func GetObjectMesh(vao uint32, vbo uint32, prog uint32) ObjectMesh{

	return ObjectMesh{
				vao: vao,
				vbo: vbo,
				prog: prog,
			}
}

func (c *ObjectMesh) Draw(rc *renderer.RenderContext, worldObject model.WorldObject){

	gl.UseProgram(c.prog)
	gl.BindVertexArray(c.vao)

	colorLoc := gl.GetUniformLocation(c.prog, gl.Str("uColor\x00"))	
	color := worldObject.GetColor()
	gl.Uniform4f(colorLoc, color.R, color.G, color.B, color.A)

	mvp := rc.OrthoProjection.	
				Mul4(rc.ViewMat).
				Mul4(worldObject.GetTransform().GetModel())

	mvpLoc := gl.GetUniformLocation(c.prog, gl.Str("uMVP\x00"))
	gl.UniformMatrix4fv(mvpLoc, 1, false, &mvp[0])

	gl.DrawArrays(gl.TRIANGLE_FAN, 0 , 4)
}
