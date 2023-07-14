package utils

// Get config path for local or docker
func GetConfigPath(configPath string) string {
	if configPath == "docker" {
		return "./configs/config-docker"
	}
	return "./configs/config-local"
}