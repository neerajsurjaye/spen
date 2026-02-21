package state

import "github.com/neerajsurjaye/spen/internal/model"

/*
Contains information related to the current running engine.
*/
type Engine struct {
	UiInfo    model.UiInfo
	GlfwState model.GlfwState
}

var engine *Engine = nil

func InitEngine() {
	if engine == nil {
		engine = &Engine{}
	}
}

func EngineInstance() *Engine {
	return engine
}