package config

// ReadDefault reads configuration file named `config` located in current folder
func ReadDefault(target interface{}) error {
	return (Source{FileName: "config"}).Read(target)
}
