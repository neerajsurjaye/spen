package mesh

import (
	"github.com/go-gl/gl/all-core/gl"
	"github.com/neerajsurjaye/spen/internal/model"
	"github.com/neerajsurjaye/spen/internal/renderer"
)

type WallMesh struct {
	vao uint32
	vbo uint32
	prog uint32
}


func GetWall(vao uint32, vbo uint32, prog uint32) WallMesh{

	return WallMesh{
				vao: vao,
				vbo: vbo,
				prog: prog,
			}
}

func (c *WallMesh) Draw(rc *renderer.RenderContext, wall *model.Wall){

	gl.UseProgram(c.prog)
	gl.BindVertexArray(c.vao)

	colorLoc := gl.GetUniformLocation(c.prog, gl.Str("uColor\x00"))	
	gl.Uniform4f(colorLoc, wall.Color.R, wall.Color.G, wall.Color.B, wall.Color.A)

	mvp := rc.OrthoProjection.	
				Mul4(rc.ViewMat).
				Mul4(wall.Transform.GetModel())

	mvpLoc := gl.GetUniformLocation(c.prog, gl.Str("uMVP\x00"))
	gl.UniformMatrix4fv(mvpLoc, 1, false, &mvp[0])

	gl.DrawArrays(gl.TRIANGLE_FAN, 0 , 4)
}
