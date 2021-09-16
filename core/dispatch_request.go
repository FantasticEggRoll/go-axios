package core

import (
	"errors"
	"go-axios/go_promise"
)

func dispatcherRequest(config *Config) (*go_promise.Promise, error) {
	if config == nil {
		return nil, errors.New("config is nil")
	}

	err := TransformRequestData(config)
	if err != nil {
		return nil, err
	}

	adapter := config.Adapter
	adapter(config).Then(func(value interface{}) (interface{}, error) {
		value, err := TransformResponseData(value, config)
		if err != nil {
			return nil, err
		}

	}, func(reason error) (interface{}, error) {

	})

	return go_promise.NewPromise(func(resolve go_promise.Resolve, reject go_promise.Reject) error {

		return nil
	}), nil
}
