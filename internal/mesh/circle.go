package mesh

import (
	"github.com/go-gl/gl/all-core/gl"
	"github.com/neerajsurjaye/spen/internal/utils"
)

type CircleInfo struct {
	circleVert []float32
	Transform CircleTransform

	vao uint32
	vbo uint32
	prog uint32

	circleColor [4]float32

}

type CircleTransform struct{
	X float64
	Y float64
	R float64
}

var circleVert []float32 = []float32{
							-1, -1,
							1, -1,
							1, 1,
							-1, 1,
						}

func GetCircle(x float64, y float64, r float64) CircleInfo{
	circleTransform := CircleTransform{X : x, Y : y , R : r}

	return CircleInfo{
				circleVert:  circleVert,
				Transform: circleTransform,
			}
}

func (c *CircleInfo) ComputeCircleVAO() {
	var vao, vbo uint32

	gl.GenVertexArrays(1, &vao)
	gl.GenBuffers(1, &vbo)

	gl.BindVertexArray(vao)

	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(
		gl.ARRAY_BUFFER,
		len(c.circleVert) * 4,
		gl.Ptr(c.circleVert),
		gl.STATIC_DRAW,
	)

	gl.VertexAttribPointer(0, 2, gl.FLOAT, false, 0, nil)
	gl.EnableVertexAttribArray(0)

	c.vao = vao
	c.vbo = vbo
} 

func (c *CircleInfo) AttachShader(fragShaderPath string, vertexShaderPath string){
	//Use gl.shaderType to match shader here
	fragShaderSrc, err := utils.ReadShader(fragShaderPath)
	if err != nil{
		panic(err)
	}
	vertexShaderSrc, err := utils.ReadShader(vertexShaderPath)
	if err != nil{
		panic(err)
	}

	fragShader, err := utils.CompileShader(fragShaderSrc, gl.FRAGMENT_SHADER)
	if err != nil{
		panic(err)
	}

	vertexShader, err := utils.CompileShader(vertexShaderSrc, gl.VERTEX_SHADER)
	if err != nil{
		panic(err)
	}

	prog := gl.CreateProgram()
	gl.AttachShader(prog, vertexShader)
	gl.AttachShader(prog, fragShader)
	gl.LinkProgram(prog)

	utils.CheckProgLinkStatus(prog)

	c.prog = prog

	gl.DeleteShader(fragShader)
	gl.DeleteShader(vertexShader)

}

func (c *CircleInfo) SetColor(r float32, g float32, b float32, a float32){
	c.circleColor = [4]float32{r, g, b, a}
}

func (c *CircleInfo) Draw(){
	gl.UseProgram(c.prog)
	gl.BindVertexArray(c.vao)

	colorLoc := gl.GetUniformLocation(c.prog, gl.Str("uColor\x00"))
	gl.Uniform4f(colorLoc, c.circleColor[0], c.circleColor[1], c.circleColor[2], c.circleColor[3])

	gl.DrawArrays(gl.TRIANGLE_FAN, 0 , 4)
}