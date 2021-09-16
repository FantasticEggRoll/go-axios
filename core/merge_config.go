package core

//MergeConfig
/**
 * @param {Config} config1
 * @param {Config} config2
 * @returns {Object} New object resulting from merging config2 to config1
 */
func MergeConfig(configs ...*Config) *Config {
	mergedConfig := &Config{}

	for _, config := range configs {
		if config.URL != "" {
			mergedConfig.URL = config.URL
		}
		if config.Method != "" {
			mergedConfig.Method = config.Method
		}
	}

	return mergedConfig
}
