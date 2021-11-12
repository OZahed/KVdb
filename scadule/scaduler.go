package scadule

import (
	"log"
	"runtime"
)

func ScaduleTask(t func() error, name string, l *log.Logger) error {
	go func() {
		for {
			if err := t(); err != nil {
				l.Panicln(name, " err: ", err)
			}
			l.Println("task: ", name, "done")
		}
	}()
	l.Println("Task, ", name, "started... (", runtime.NumGoroutine(), ") goroutines are running")
	return nil
}
