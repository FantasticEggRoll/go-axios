package core

import "go-axios/go_promise"

type handler struct {
	fulfilled   go_promise.OnFulFilled
	rejected    go_promise.OnRejected
	synchronous bool
}

type InterceptorManager struct {
	handlers []*handler
}

func NewInterceptorManager() *InterceptorManager {
	return &InterceptorManager{
		handlers: make([]*handler, 0),
	}
}

func (im *InterceptorManager) Use(onFulfilled go_promise.OnFulFilled, onRejected go_promise.OnRejected, options interface{}) int {
	im.handlers = append(im.handlers, &handler{
		fulfilled: onFulfilled,
		rejected:  onRejected,
	})
	return len(im.handlers) - 1
}

func (im *InterceptorManager) Eject(id int) bool {
	if id >= len(im.handlers) || id < 0 {
		return false
	}
	if im.handlers[id] != nil {
		im.handlers[id] = nil
		return true
	}
	return false
}

func (im *InterceptorManager) ForEach() {

}
