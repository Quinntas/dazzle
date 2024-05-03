package web

import "context"

type HttpRouter struct {
	Path string
	ctx  context.Context
}

func NewHttpRouter(path string, ctx context.Context) HttpRouter {
	return HttpRouter{
		Path: path,
		ctx:  ctx,
	}
}

type Router func(path string, ctx context.Context)
