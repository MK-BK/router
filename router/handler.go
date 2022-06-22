package router

import (
	"io"
	"net/http"
)

type Handler http.HandlerFunc

func (r *Router) HandlerFunc(method, path string, handle Handler) {
	if r.tree[method] == nil {
		r.tree[method] = NewTreeNode()
	}
	tree := r.tree[method]
	tree.Insert(path, handle)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.Handler(w, req)
	return
}

func (r *Router) Handler(w http.ResponseWriter, req *http.Request) {
	method := req.Method
	path := req.URL.Path

	tree, ok := r.tree[method]
	if !ok {
		io.WriteString(w, "handler not found")
		return
	}
	handler, err := tree.Search(path)
	if err != nil {
		io.WriteString(w, "handler not found")
		return
	}

	handlers := r.middles
	handlers = append(handlers, handler)

	for i := len(handlers) - 1; i >= 0; i-- {
		handlers[i](w, req)
	}
}
