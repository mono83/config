package config

import "os"

// pathConcat concatenates path and file
func pathConcat(path, file string) string {
	if path[len(path)-1] != os.PathSeparator {
		return path + string(os.PathSeparator) + file
	}

	return path + file
}
