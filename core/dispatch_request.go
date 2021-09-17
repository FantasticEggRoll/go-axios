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
	return adapter(config).Then(func(response interface{}) (interface{}, error) {
		response, err := TransformResponseData(response, config)
		return response, err
	}, func(reason interface{}) (interface{}, error) {
		reason, err := TransformResponseData(reason, config)
		if err != nil {
			return nil, err
		}
		return go_promise.RejectPromise(reason), nil
	}), nil
}
