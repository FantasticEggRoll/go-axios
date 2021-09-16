package core

import (
	"errors"
	"fmt"
)

func TransformRequestData(config *Config) error {
	for _, transformer := range config.TransformRequests {
		data, err := transformer.Transform(config.Data, config.Header)
		if err != nil {
			return errors.New(fmt.Sprintf("transform request data error, %v", err))
		}
		config.Data = data
	}
	return nil
}

func TransformResponseData(value interface{}, config *Config) (interface{}, error) {
	var err error
	for _, transformer := range config.TransformerResponse {
		value, err = transformer.Transform(value)
		if err != nil {
			return value, errors.New(fmt.Sprintf("transform request data error, %v", err))
		}
	}
	return value, nil
}
