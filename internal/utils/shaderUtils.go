package utils

import (
	"fmt"
	"strings"

	"github.com/go-gl/gl/all-core/gl"
)

func CompileShader(source string, shaderType uint32) (uint32, error){
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)

	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("Failed to compil %v: %v", source, log)
	}

	return shader, nil
}

func CreateProgram(fragShaderPath string, vertexShaderPath string) uint32{
	//Use gl.shaderType to match shader here
	fragShaderSrc, err := ReadShader(fragShaderPath)
	if err != nil{
		panic(err)
	}
	vertexShaderSrc, err := ReadShader(vertexShaderPath)
	if err != nil{
		panic(err)
	}

	fragShader, err := CompileShader(fragShaderSrc, gl.FRAGMENT_SHADER)
	if err != nil{
		panic(err)
	}

	vertexShader, err := CompileShader(vertexShaderSrc, gl.VERTEX_SHADER)
	if err != nil{
		panic(err)
	}

	prog := gl.CreateProgram()
	gl.AttachShader(prog, vertexShader)
	gl.AttachShader(prog, fragShader)
	gl.LinkProgram(prog)

	CheckProgLinkStatus(prog)

	gl.DeleteShader(fragShader)
	gl.DeleteShader(vertexShader)

	return prog
}