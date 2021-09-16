package helpers

import (
	"encoding/json"
	"errors"
	go_axios "go-axios"
	"go-axios/core"
	neturl "net/url"
	"strings"
)

//BuildURL
/**
 * Build a URL by appending params to the end
 *
 * @param {string} url The base of the url (e.g., http://www.google.com)
 * @param {object} [params] The params to be appended
 * @returns {error} error
 * @returns {string} The formatted url
 */
func BuildURL(url string, param interface{}, serializeParam core.ParamSerializer) (string, error) {
	if param == nil {
		return url, nil
	}

	var serialized string
	var err error
	if serializeParam != nil {
		serialized, err = serializeParam.Serialize(param)
		if err != nil {
			return "", err
		}
	} else if tmpParam, ok := go_axios.IsURLSearchParams(param); ok {
		serialized = tmpParam.Encode()
	} else {
		paramMap, ok := param.(map[string]interface{})
		if !ok {
			return "", errors.New("can not parse param")
		}
		serialized, err = serialize(paramMap)
		if err != nil {
			return "", err
		}
	}

	if serialized != "" {
		index := strings.Index(url, "#")
		if index != -1 {
			url = url[0:index]
		}
		index = strings.Index(url, "?")
		if index != -1 {
			url += "&"
		} else {
			url += "?"
		}
		url += serialized
	}

	return url, nil
}

func serialize(param map[string]interface{}) (string, error) {
	parts := make([]string, 0)
	for key, value := range param {
		if value == nil {
			continue
		}
		if go_axios.IsArray(key) || go_axios.IsSlice(key) {
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
