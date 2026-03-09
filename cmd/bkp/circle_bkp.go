// package mesh

// import (
// 	"github.com/go-gl/gl/all-core/gl"
// 	"github.com/go-gl/mathgl/mgl32"
// 	"github.com/neerajsurjaye/spen/internal/renderer"
// )

// type CircleMesh struct {
// 	CircleVert []float32

// 	vao uint32
// 	vbo uint32
// 	prog uint32

// 	circleColor [4]float32
// 	Transform CircleTransform
// }

// type CircleTransform struct{
// 	X float64
// 	Y float64
// 	R float64
// }

// func GetCircle(vao uint32, vbo uint32, prog uint32) CircleMesh{

// 	return CircleMesh{
// 				vao: vao,
// 				vbo: vbo,
// 				prog: prog,
// 			}
// }

// func (c *CircleMesh) SetColor(r float32, g float32, b float32, a float32){
// 	c.circleColor = [4]float32{r, g, b, a}
// }

// func (c *CircleMesh) SetTransform(x float64, y float64, r float64){
// 	circleTransform := CircleTransform{X: x, Y: y, R: r}
// 	c.Transform = circleTransform
// }

// func (c *CircleMesh) getTranslation() mgl32.Mat4{
// 	return mgl32.Translate3D(float32(c.Transform.X), float32(c.Transform.Y), 0);
// }

// func (c *CircleMesh) getScale() mgl32.Mat4{
// 	return mgl32.Scale3D(float32(c.Transform.R), float32(c.Transform.R), 1)
// }

// func (c *CircleMesh) getModel() mgl32.Mat4{
// 	return c.getTranslation().Mul4(c.getScale())
// }

// func (c *CircleMesh) Draw(rc *renderer.RenderContext){

// 	// if(rc.DebugDraw){
// 	// 	aabb := AABB{}
// 	// }

// 	gl.UseProgram(c.prog)
// 	gl.BindVertexArray(c.vao)

// 	colorLoc := gl.GetUniformLocation(c.prog, gl.Str("uColor\x00"))
// 	gl.Uniform4f(colorLoc, c.circleColor[0], c.circleColor[1], c.circleColor[2], c.circleColor[3])

// 	mvp := rc.OrthoProjection.	
// 				Mul4(rc.ViewMat).
// 				Mul4(c.getModel())

// 	mvpLoc := gl.GetUniformLocation(c.prog, gl.Str("uMVP\x00"))
// 	gl.UniformMatrix4fv(mvpLoc, 1, false, &mvp[0])

// 	gl.DrawArrays(gl.TRIANGLE_FAN, 0 , 4)
// }
