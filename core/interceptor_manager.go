package core

import "go-axios/go_promise"

type Handler struct {
	OnFulfilled go_promise.OnFulFilled
	OnRejected  go_promise.OnRejected
	synchronous bool
}

type InterceptorManager struct {
	handlers []*Handler
}

func NewHandler(onFulFilled go_promise.OnFulFilled, onRejected go_promise.OnRejected, synchronous bool) *Handler {
	return &Handler{
		OnFulfilled: onFulFilled,
		OnRejected:  onRejected,
		synchronous: synchronous,
	}
}

func NewInterceptorManager() *InterceptorManager {
	return &InterceptorManager{
		handlers: make([]*Handler, 0),
	}
}

func (im *InterceptorManager) Use(onFulfilled go_promise.OnFulFilled, onRejected go_promise.OnRejected, options interface{}) int {
	im.handlers = append(im.handlers, &Handler{
		OnFulfilled: onFulfilled,
		OnRejected:  onRejected,
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
