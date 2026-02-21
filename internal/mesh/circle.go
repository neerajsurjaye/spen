package mesh

import (
	"github.com/go-gl/gl/all-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type CircleInfo struct {
	CircleVert []float32

	vao uint32
	vbo uint32
	prog uint32

	circleColor [4]float32
	Transform CircleTransform
}

type CircleTransform struct{
	X float64
	Y float64
	R float64
}

func GetCircle(vao uint32, vbo uint32, prog uint32) CircleInfo{

	return CircleInfo{
				vao: vao,
				vbo: vbo,
				prog: prog,
			}
}

func (c *CircleInfo) SetColor(r float32, g float32, b float32, a float32){
	c.circleColor = [4]float32{r, g, b, a}
}

func (c *CircleInfo) SetTransform(x float64, y float64, r float64){
	circleTransform := CircleTransform{X: x, Y: y, R: r}
	c.Transform = circleTransform
}

func (c *CircleInfo) GetTranslation() mgl32.Mat4{
	return mgl32.Translate3D(float32(c.Transform.X), float32(c.Transform.Y), 0);
}

func (c *CircleInfo) Draw(projection mgl32.Mat4, view mgl32.Mat4){

	gl.UseProgram(c.prog)
	gl.BindVertexArray(c.vao)

	colorLoc := gl.GetUniformLocation(c.prog, gl.Str("uColor\x00"))
	gl.Uniform4f(colorLoc, c.circleColor[0], c.circleColor[1], c.circleColor[2], c.circleColor[3])

	mvp := projection.
				Mul4(view).
				Mul4(c.GetTranslation())

	mvpLoc := gl.GetUniformLocation(c.prog, gl.Str("uMVP\x00"))
	gl.UniformMatrix4fv(mvpLoc, 1, false, &mvp[0])

	radiusLoc := gl.GetUniformLocation(c.prog, gl.Str("uRadius\x00"))
	gl.Uniform1f(radiusLoc, float32(c.Transform.R))

	gl.DrawArrays(gl.TRIANGLE_FAN, 0 , 4)
}