package go_promise

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestPromiseAsync(t *testing.T) {
	promise1 := NewPromise(func(resolve Resolve, reject Reject) error {
		fmt.Println("start")
		c := make(chan int, 0)
		go func() {
			time.Sleep(2 * time.Second)
			c <- 1
		}()

		<-c
		resolve("yes")
		//reject(errors.New("oh my god"))
		return nil
	})

	promise1.Then(func(value interface{}) (interface{}, error) {
		fmt.Println("haha 1", value)
		return nil, nil
	}, func(reason interface{}) (interface{}, error) {
		fmt.Println("bye 1", reason)
		return nil, nil
	})
}

func TestPromiseChain(t *testing.T) {
	promise1 := NewPromise(func(resolve Resolve, reject Reject) error {
		resolve("yes")
		//reject(errors.New("oh my god"))
		return nil
	})

	promise1.Then(func(value interface{}) (interface{}, error) {
		fmt.Println("success 1", value)
		return value, errors.New("ops")
	}, func(reason interface{}) (interface{}, error) {
		fmt.Println("fail 1", reason)
		err, _ := reason.(error)
		return nil, err
	}).Then(func(value interface{}) (interface{}, error) {
		fmt.Println("success 2", value)
		return nil, nil
	}, func(reason interface{}) (interface{}, error) {
		fmt.Println("fail 2", reason)
		err, _ := reason.(error)
		return nil, err
	}).Then(func(value interface{}) (interface{}, error) {
		fmt.Println("success 3", value)
		return value, nil
	}, func(reason interface{}) (interface{}, error) {
		fmt.Println("try to recover")
		promise2 := NewPromise(func(resolve Resolve, reject Reject) error {
			resolve("recover success")
			return nil
		})
		return promise2, nil
	}).Then(func(value interface{}) (interface{}, error) {
		fmt.Println("success 4", value)
		promise3 := NewPromise(func(resolve Resolve, reject Reject) error {
			reject(errors.New("fail again"))
			return nil
		})
		return promise3, nil
	}, func(reason interface{}) (interface{}, error) {
		fmt.Println("fail 4")
		return nil, nil
	}).Then(func(value interface{}) (interface{}, error) {
		fmt.Println("success 5", value)
		return value, nil
	}, func(reason interface{}) (interface{}, error) {
		fmt.Println("fail 5", reason)
		return "god damn", nil
	}).Then(func(value interface{}) (interface{}, error) {
		time.Sleep(1 * time.Second)
		fmt.Println("success 6", value)
		return value, nil
	}, func(reason interface{}) (interface{}, error) {
		fmt.Println("fail 6", reason)
		return nil, nil
	})
}
