package mesh

import (
	"github.com/go-gl/gl/all-core/gl"
	"github.com/neerajsurjaye/spen/internal/model"
	"github.com/neerajsurjaye/spen/internal/renderer"
)

type CircleMesh struct {
	vao uint32
	vbo uint32
	prog uint32
}


func GetCircle(vao uint32, vbo uint32, prog uint32) CircleMesh{

	return CircleMesh{
				vao: vao,
				vbo: vbo,
				prog: prog,
			}
}

func (c *CircleMesh) Draw(rc *renderer.RenderContext, circle *model.Circle){

	gl.UseProgram(c.prog)
	gl.BindVertexArray(c.vao)

	colorLoc := gl.GetUniformLocation(c.prog, gl.Str("uColor\x00"))	
	gl.Uniform4f(colorLoc, circle.Color.R, circle.Color.G, circle.Color.B, circle.Color.A)

	mvp := rc.OrthoProjection.	
				Mul4(rc.ViewMat).
				Mul4(circle.GetModel())

	mvpLoc := gl.GetUniformLocation(c.prog, gl.Str("uMVP\x00"))
	gl.UniformMatrix4fv(mvpLoc, 1, false, &mvp[0])

	gl.DrawArrays(gl.TRIANGLE_FAN, 0 , 4)
}
