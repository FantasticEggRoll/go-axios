package go_promise

import (
	"errors"
)

type (
	Status      int
	Resolve     func(value interface{})
	Reject      func(err interface{})
	Executor    func(resolve Resolve, reject Reject) error
	OnFulFilled func(value interface{}) (interface{}, error)
	OnRejected  func(reason interface{}) (interface{}, error)
)

const (
	PENDING Status = iota
	RESOLVED
	REJECTED
)

type Promise struct {
	value               interface{}
	reason              interface{}
	status              Status
	OnResolvedCallbacks []func()
	OnRejectedCallbacks []func()
}

func NewPromise(executor Executor) *Promise {
	promise := &Promise{
		OnResolvedCallbacks: make([]func(), 0),
		OnRejectedCallbacks: make([]func(), 0),
	}

	resolve := func(value interface{}) {
		if promise.status == PENDING {
			promise.status = RESOLVED
			promise.value = value
			for _, callback := range promise.OnResolvedCallbacks {
				callback()
			}
		}
	}

	reject := func(reason interface{}) {
		if promise.status == PENDING {
			promise.status = REJECTED
			promise.reason = reason
			for _, callback := range promise.OnRejectedCallbacks {
				callback()
			}
		}
	}

	err := executor(resolve, reject)
	if err != nil {
		reject(err)
	}

	return promise
}

func (promise *Promise) Then(onFullFilled OnFulFilled, onRejected ...OnRejected) *Promise {
	var promise2 *Promise
	promise2 = NewPromise(func(resolve Resolve, reject Reject) error {
		if promise.status == RESOLVED {
			result, err := onFullFilled(promise.value)
			if err != nil {
				reject(err)
				return err
			}
			resolvePromise(promise2, result, resolve, reject)
		}
		if promise.status == REJECTED {
			result, err := onRejected[0](promise.reason)
			if err != nil {
				reject(err)
				return err
			}
			resolvePromise(promise2, result, resolve, reject)
		}
		if promise.status == PENDING {
			promise.OnResolvedCallbacks = append(promise.OnResolvedCallbacks, func() {
				result, err := onFullFilled(promise.value)
				if err != nil {
					reject(err)
					return
				}
				resolvePromise(promise2, result, resolve, reject)
			})
			promise.OnRejectedCallbacks = append(promise.OnRejectedCallbacks, func() {
				result, err := onRejected[0](promise.reason)
				if err != nil {
					reject(err)
					return
				}
				resolvePromise(promise2, result, resolve, reject)
			})
		}
		return nil
	})
	return promise2
}

func (promise *Promise) Catch(onRejected OnRejected) *Promise {
	return promise.Then(nil, onRejected)
}

func ResolvePromise(value interface{}) *Promise {
	return NewPromise(func(resolve Resolve, reject Reject) error {
		resolve(value)
		return nil
	})
}

func RejectPromise(reason interface{}) *Promise {
	return NewPromise(func(resolve Resolve, reject Reject) error {
		reject(reason)
		return nil
	})
}

func (promise *Promise) All(promises []*Promise) *Promise {
	return NewPromise(func(resolve Resolve, reject Reject) error {
		length := len(promises)
		arr := make([]interface{}, length)
		i := 0
		processData := func(index int, data interface{}) {
			arr[index] = data
			i += 1
			if i == length {
				resolve(arr)
			}
		}

		for index, promise := range promises {
			promise.Then(func(value interface{}) (interface{}, error) {
				processData(index, value)
				return nil, nil
			}, func(reason interface{}) (interface{}, error) {
				reject(reason)
				return nil, nil
			})
		}

		return nil
	})
}

func resolvePromise(promise *Promise, x interface{}, resolve Resolve, reject Reject) {
	if promise == x {
		reject(errors.New("chaining cycle detected for go_promise"))
	}
	called := false
	if tmpPromise, ok := x.(*Promise); ok {
		tmpPromise.Then(func(y interface{}) (interface{}, error) {
			if called {
				return nil, nil
			}
			called = true
			resolvePromise(promise, y, resolve, reject)
			return nil, nil
		}, func(err interface{}) (interface{}, error) {
			if called {
				return nil, nil
			}
			called = true
			reject(err)
			return nil, nil
		})
	} else {
		resolve(x)
	}
}
