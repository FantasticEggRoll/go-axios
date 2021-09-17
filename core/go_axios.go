package core

import (
	"go-axios/go_promise"
	"strings"
)

type Interceptors struct {
	request  *InterceptorManager
	response *InterceptorManager
}

type Axios struct {
	defaults     *Config
	interceptors *Interceptors
}

func Create(instanceConfig *Config) *Axios {
	return &Axios{
		defaults: instanceConfig,
		interceptors: &Interceptors{
			NewInterceptorManager(),
			NewInterceptorManager(),
		},
	}
}

func (axios *Axios) Request(config *Config) *go_promise.Promise {
	//config = MergeConfig(axios.defaults, config)

	requestInterceptorChain := make([]interface{}, 0)
	//头插构造请求拦截链
	for _, handler := range axios.interceptors.request.handlers {
		pair := []interface{}{handler.fulfilled, handler.rejected}
		requestInterceptorChain = append(pair, requestInterceptorChain)
	}

	responseInterceptorChain := make([]interface{}, 0)
	//尾插构造响应拦截链
	synchronousRequestInterceptors := true
	for _, handler := range axios.interceptors.response.handlers {
		synchronousRequestInterceptors = synchronousRequestInterceptors && handler.synchronous
		responseInterceptorChain = append(responseInterceptorChain, handler.fulfilled, handler.rejected)
	}

	if !synchronousRequestInterceptors {
		chain := []interface{}{dispatcherRequest, nil}
		chain = append(requestInterceptorChain, chain, responseInterceptorChain)
		promise := go_promise.ResolvePromise(config)
		for i := 0; i < len(chain); i += 2 {
			onFulFilled, _ := chain[i].(go_promise.OnFulFilled)
			onRejected, _ := chain[i+1].(go_promise.OnRejected)
			promise = promise.Then(onFulFilled, onRejected)
		}

		return promise
	}

	for i := 0; i < len(requestInterceptorChain); i += 2 {
		onFulFilled, _ := requestInterceptorChain[i].(go_promise.OnFulFilled)
		onRejected, _ := requestInterceptorChain[i+1].(go_promise.OnRejected)
		result, err := onFulFilled(config)
		config, _ = result.(*Config)
		if err != nil {
			onRejected(err)
			break
		}
	}

	promise, err := dispatcherRequest(config)
	if err != nil {
		return go_promise.RejectPromise(err)
	}

	for i := 0; i < len(responseInterceptorChain); i += 2 {
		onFulFilled, _ := responseInterceptorChain[i].(go_promise.OnFulFilled)
		onRejected, _ := responseInterceptorChain[i+1].(go_promise.OnRejected)
		promise = promise.Then(onFulFilled, onRejected)
	}

	return promise
}

func (axios *Axios) RequestWithoutData(url string, method Method, config *Config) *go_promise.Promise {
	/*
		MergeConfig(axios.defaults, &Config{
			URL:    url,
			Method: method,
			//Data:   config.Data,
		})

	*/
	return axios.Request(axios.defaults)
}

func (axios *Axios) RequestWithData(url string, method Method, data interface{}, config *Config) *go_promise.Promise {
	MergeConfig(axios.defaults, &Config{
		URL:    url,
		Method: method,
		Data:   data,
	})
	return axios.Request(config)
}

func (axios *Axios) Get(url string, config *Config) *go_promise.Promise {
	return axios.RequestWithoutData(url, GET, config)
}

func (axios *Axios) Delete(url string, config *Config) *go_promise.Promise {
	return axios.RequestWithoutData(url, DELETE, config)
}

func (axios *Axios) Head(url string, config *Config) *go_promise.Promise {
	return axios.RequestWithoutData(url, HEAD, config)
}

func (axios *Axios) Options(url string, config *Config) *go_promise.Promise {
	return axios.RequestWithoutData(url, OPTIONS, config)
}

func (axios *Axios) Post(url string, data interface{}, config *Config) *go_promise.Promise {
	return axios.RequestWithData(url, POST, data, config)
}

func (axios *Axios) Put(url string, data interface{}, config *Config) *go_promise.Promise {
	return axios.RequestWithData(url, PUT, data, config)
}

func (axios *Axios) Patch(url string, data interface{}, config *Config) *go_promise.Promise {
	return axios.RequestWithData(url, PATCH, data, config)
}

func (axios *Axios) GetURL(config *Config) (string, error) {
	config = MergeConfig(axios.defaults, config)
	url, err := BuildURL(config.URL, config.Param, config.SerializeParam)
	if err != nil {
		return "", err
	}
	return strings.ReplaceAll(url, `\`, ""), nil
}
