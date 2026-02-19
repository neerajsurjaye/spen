package logger

import (
	"fmt"
	"os"
)

type Logger struct {
	ch   chan string
	done chan struct{}
}

func (l *Logger) run() {
	for {
		select {
		case msg, ok := <-l.ch:
			if !ok{
				return
			}
			fmt.Fprintln(os.Stdout, msg)
		case <-l.done:
			//close(done) will wake this line up
			for {
				select{
				case msg := <- l.ch:
						fmt.Fprintln(os.Stdout, msg)
				default:
					return
				}
			
			}
		}
	}
}

func (l *Logger) Log(msg ...any){
	select{
	/**
	If can push to channe run this. If channe is full go to default case
	*/
	case l.ch <- fmt.Sprint(msg...):
	default:
	}
}

func (l *Logger) Close(){
	close(l.done)
}

func NewLogger(buffer int) *Logger {
	l := &Logger{
		ch:   make(chan string, buffer),
		done: make(chan struct{}),
	}

	go l.run()
	return l
}
