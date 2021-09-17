package core

//MergeConfig
/**
 * @param {Config} config1
 * @param {Config} config2
 * @returns {Object} New object resulting from merging config2 to config1
 */
func MergeConfig(configs ...*Config) *Config {
	mergedConfig := NewConfig()

	for _, config := range configs {
		if config == nil {
			continue
		}
		if config.URL != "" {
			mergedConfig.URL = config.URL
		}
		if config.Method != "" {
			mergedConfig.Method = config.Method
		}
		if config.Header != nil {
			for k, vs := range config.Header.Header {
				for _, v := range vs {
					mergedConfig.Header.Add(k, v)

				}
			}
		}
		if config.Param != nil {
			for k, vs := range config.Param.Values {
				for _, v := range vs {
					mergedConfig.Param.Add(k, v)

				}
			}
		}
		if config.SerializeParam != nil {
			mergedConfig.SerializeParam = config.SerializeParam
		}
		if config.Data != nil {
			mergedConfig.Data = config.Data
		}
		if config.TransformRequests != nil {
			mergedConfig.TransformRequests = config.TransformRequests
		}
		if config.TransformerResponse != nil {
			mergedConfig.TransformerResponse = config.TransformerResponse
		}
		if config.Adapter != nil {
			mergedConfig.Adapter = config.Adapter
		}
	}

	return mergedConfig
}
