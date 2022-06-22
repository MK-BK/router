package router

import (
	"fmt"
	"net/http"
)

type Logger struct {
	level int
}

func (l *Logger) Info(values ...interface{}) {
	fmt.Println(values...)
}

func (l *Logger) Error(value interface{}) {

}

func NewLogger() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log := &Logger{}
		log.Info(r.Method, r.URL.Path)
	}

}
