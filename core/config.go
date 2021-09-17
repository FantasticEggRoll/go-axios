package core

import (
	"go-axios/go_promise"
	"net/http"
	"net/url"
)

type (
	Method string
)

type Header struct {
	http.Header
}

type Param struct {
	url.Values
}

const (
	GET     Method = http.MethodGet
	DELETE  Method = http.MethodDelete
	HEAD    Method = http.MethodHead
	OPTIONS Method = http.MethodOptions

	POST  Method = http.MethodPost
	PUT   Method = http.MethodPut
	PATCH Method = http.MethodPatch
)

type ParamSerializer interface {
	Serialize(param interface{}) (string, error)
}

type SerializeParam func(param interface{}) (string, error)

func (serializer SerializeParam) Serialize(param interface{}) (string, error) {
	return serializer(param)
}

// RequestTransformer

type RequestTransformer interface {
	Transform(interface{}, Header) (interface{}, error)
}

type TransformerRequest func(interface{}, Header) (interface{}, error)

func (req TransformerRequest) Transform(data interface{}, header Header) (interface{}, error) {
	return req(data, header)
}

// ResponseTransformer

type ResponseTransformer interface {
	Transform(interface{}) (interface{}, error)
}

type TransformerResponse func(interface{}) (interface{}, error)

func (resp TransformerResponse) Transform(data interface{}) (interface{}, error) {
	return resp(data)
}

type Adapter func(config *Config) *go_promise.Promise

type Config struct {
	URL                 string `json:"url"`
	Method              Method `json:"method"`
	Header              Header `json:"header"`
	Param               Param  `json:"param"`
	SerializeParam      ParamSerializer
	Data                interface{}
	TransformRequests   []RequestTransformer
	TransformerResponse []ResponseTransformer
	Adapter             Adapter
}

func NewConfig() *Config {
	return &Config{
		URL:                 "",
		Method:              GET,
		Header:              NewHeader(),
		Param:               NewParam(),
		Data:                nil,
		SerializeParam:      SerializeParam(DefaultSerializeParam),
		TransformRequests:   []RequestTransformer{TransformerRequest(DefaultTransformRequest)},
		TransformerResponse: []ResponseTransformer{TransformerResponse(DefaultTransformResponse)},
		Adapter:             DefaultAdapter,
	}
}

func (config *Config) AddRequestTransform(transformer RequestTransformer) {
	config.TransformRequests = append(config.TransformRequests, transformer)
}

func NewParam() Param {
	return Param{
		url.Values{},
	}
}

func NewHeader() Header {
	return Header{
		http.Header{},
	}
}
