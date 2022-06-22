package router

type Context struct {
	Params map[string]interface{}
}

func (c *Context) Get(key string) (interface{}, bool) {
	value, ok := c.Params[key]
	return value, ok
}
