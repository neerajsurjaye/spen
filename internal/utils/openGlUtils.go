package utils

import (
	"strings"

	"github.com/go-gl/gl/all-core/gl"
)

func CheckProgLinkStatus(prog uint32) {
	var status int32
	gl.GetProgramiv(prog, gl.LINK_STATUS, &status)

	if status == gl.FALSE{
		var logLength int32
		gl.GetProgramiv(prog, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength + 1))
		gl.GetProgramInfoLog(prog, logLength, nil, gl.Str(log))

		panic("Program link failed : " + log)
	}
}