package mesh

import (
	"github.com/go-gl/gl/all-core/gl"
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

func (c *CircleInfo) Draw(){
	gl.UseProgram(c.prog)
	gl.BindVertexArray(c.vao)

	colorLoc := gl.GetUniformLocation(c.prog, gl.Str("uColor\x00"))
	gl.Uniform4f(colorLoc, c.circleColor[0], c.circleColor[1], c.circleColor[2], c.circleColor[3])

	gl.DrawArrays(gl.TRIANGLE_FAN, 0 , 4)
}