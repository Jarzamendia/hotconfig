package hotconfig

import "os"

//GetEnvVarOrDefault Get a env variable, if null, a default value.
func GetEnvVarOrDefault(key string, defaultValue string) string {

	result := os.Getenv(key)

	if result == "" {
		return defaultValue
	}

	return result
}
