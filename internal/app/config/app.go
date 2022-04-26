package config

import "os"

var configPath = "./config.yaml"

func isFile(fp string) bool {
	f, e := os.Stat(fp)
	if e != nil {
		return false
	}

	return !f.IsDir()
}

func SetConfigPath(path string) {
	if !isFile(path) {
		panic("config filepath does not exist")
	}

	configPath = path
}
