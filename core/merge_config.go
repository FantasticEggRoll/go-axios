package core

//MergeConfig
/**
 * @param {Config} config1
 * @param {Config} config2
 * @returns {Object} New object resulting from merging config2 to config1
 */
func MergeConfig(configs ...*Config) *Config {
	mergedConfig := Config{}

	for _, config := range configs {
		if config == nil {
			continue
		}
		mergedConfig.URL = config.URL
		mergedConfig.Method = config.Method
		mergedConfig.Header = config.Header
		mergedConfig.Param = config.Param
		mergedConfig.SerializeParam = config.SerializeParam
		mergedConfig.Data = config.Data
		mergedConfig.TransformRequests = config.TransformRequests
		mergedConfig.TransformerResponse = config.TransformerResponse
		mergedConfig.Adapter = config.Adapter
	}

	return &mergedConfig
}
