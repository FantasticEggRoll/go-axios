package core

import (
	"fmt"
	"testing"
)

func TestAxios(t *testing.T) {

	config := NewConfig()
	config.URL = "http://localhost:8000/hello"
	config.Method = GET
	axios := Create(config)
	axios.Get("http://localhost:8000/hello", nil).Then(func(value interface{}) (interface{}, error) {
		fmt.Println(value)
		return nil, nil
	}, func(reason interface{}) (interface{}, error) {
		fmt.Println(reason)
		return nil, nil
	})
	/*
		c := http.Client{}
		req, _ := http.NewRequest(http.MethodGet, "http://localhost:8000/hello", nil)
		rsp, err := c.Do(req)
		defer rsp.Body.Close()
		data, err := ioutil.ReadAll(rsp.Body)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(data)

	*/
}
