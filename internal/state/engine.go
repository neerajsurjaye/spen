package state

import (
	"github.com/neerajsurjaye/spen/internal/mesh"
	"github.com/neerajsurjaye/spen/internal/model"
)

type Engine struct {
	UiInfo model.UiInfo
	GlfwState model.GlfwState
	
	//temp
	Circle mesh.CircleInfo
}

var engine *Engine = nil;

func InitEngine(){
	if engine == nil{
		engine = &Engine{}
	}
}

func EngineInstance() *Engine{
	return engine
}