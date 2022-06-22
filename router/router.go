package router

import (
	"net/http"
)

type Router struct {
	tree    map[string]*TreeNode
	middles []Handler
}

func NewRouter() *Router {
	router := &Router{
		tree: make(map[string]*TreeNode),
	}
	router.Default()
	return router
}

func (r *Router) Get(path string, handle Handler) {
	r.HandlerFunc(GET, path, handle)
}

func (r *Router) Post(path string, handle Handler) {
	r.HandlerFunc(POST, path, handle)
}

func (r *Router) Put(path string, handle Handler) {
	r.HandlerFunc(PUT, path, handle)
}

func (r *Router) Delete(path string, handle Handler) {
	r.HandlerFunc(DELETE, path, handle)
}

func (r *Router) Run() {
	http.ListenAndServe(":10086", r)
}
