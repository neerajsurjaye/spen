package state

import "github.com/go-gl/glfw/v3.3/glfw"

type InputState struct {
	PrevAction map[glfw.Key]glfw.Action
}

var inputState *InputState = nil

func InitInputState() {
	if inputState == nil {
		inputState = &InputState{
			PrevAction: make(map[glfw.Key]glfw.Action),
		}
	}
}

func InputStateInstance() *InputState {
	return inputState
}