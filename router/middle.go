package router

type middleware struct {
	handler Handler
	next    *middleware
}

func (r *Router) Use(middle Handler) {
	r.middles = append(r.middles, middle)
}

func (r *Router) Default() {
	r.middles = append(r.middles, Handler(NewLogger()))
}
