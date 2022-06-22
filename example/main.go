package main

import (
	"go-router/router"
	"io"
	"net/http"
)

func main() {
	r := router.NewRouter()

	r.Get("/hello/{id}/name", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello world")
	})

	r.Run()
}
