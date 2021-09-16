package go_axios

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"go-axios/core"
	"go-axios/go_promise"
	"go-axios/helpers"
	"io"
	"io/ioutil"
	"net/http"
	neturl "net/url"
	"strings"
)

var client *http.Client

func init() {
	client = DefaultHttpClient()
}

func DefaultHttpClient() *http.Client {
	return &http.Client{}
}

func DefaultTransformRequest(data interface{}, header core.Header) (interface{}, error) {
	helpers.NormalizeHeaderName(header, "Accept")
	helpers.NormalizeHeaderName(header, "Content-Type")

	contentType := strings.ToLower(header.Get("Content-Type"))
	switch contentType {
	case "application/json":
		b, err := json.Marshal(data)
		if err != nil {
			return nil, errors.New("transform request error: " + err.Error())
		}
		return bytes.NewBuffer(b), err
	case "application/x-www-form-urlencoded":
		param, ok := IsURLSearchParams(data)
		if !ok {
			return nil, errors.New(fmt.Sprintf("transform request error: data %v is not url search param", data))
		}
		header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
		return strings.NewReader(param.Encode()), nil
	case "multipart/form-data":

	default:
		return data, nil //可能某些情况这里没处理到，把原本的数据传出去，让后续的转换器处理
	}
	return nil, nil
}

func DefaultTransformResponse(data interface{}) (interface{}, error) {
	return data, nil
}

func DefaultSerializeParam(param interface{}) (string, error) {
	paramMap, ok := param.(map[string]interface{})
	if !ok {
		return "", errors.New("param is not map")
	}

	parts := make([]string, 0)
	for key, value := range paramMap {
		if value == nil {
			continue
		}
		if IsArray(key) || IsSlice(key) {
			key += "[]"
		} else {
			value = []interface{}{value}
		}
		value, err := json.Marshal(value)
		if err != nil {
			return "", err
		}

		query := key + "=" + string(value)
		parts = append(parts, query)
	}

	return neturl.QueryEscape(strings.Join(parts, "&")), nil
}

func DefaultAdapter(config *core.Config) *go_promise.Promise {
	return go_promise.NewPromise(func(resolve go_promise.Resolve, reject go_promise.Reject) error {
		var ioReader io.Reader
		var ok bool
		if config.Data != nil {
			ioReader, ok = config.Data.(io.Reader)
			if !ok {
				reject(errors.New("data error"))
				return nil
			}
		}

		request, err := http.NewRequest(string(config.Method), config.URL, ioReader)
		if err != nil {
			reject(err)
			return nil
		}
		resp, err := client.Do(request)
		defer func() {
			err := resp.Body.Close()
			if err != nil {
				reject(err)
			}
		}()

		if err != nil {
			reject(err)
			return nil
		}

		if resp.StatusCode != http.StatusOK {
			reject(err)
			return nil
		}

		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			reject(err)
			return nil
		}

		resolve(data)
		return nil
	})
}
