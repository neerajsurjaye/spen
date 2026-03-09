package state

import "github.com/neerajsurjaye/spen/internal/model"

/*
Contains information related to the current running engine.
*/
type Engine struct {
	UiInfo    model.UiInfo
	GlfwState model.GlfwState

	IsPaused bool
	StepSim bool
}

var engine *Engine = nil

func InitEngine() {
	if engine == nil {
		engine = &Engine{
			IsPaused: false,
		}
	}
}

func EngineInstance() *Engine {
	return engine
}