package data

import ()

//Config struct
type Config struct {
}

// Get function
func (x Config) Get(key string) (value string) {
	values := Config.List()
	value = values[key]
	return value
}

// Set function
func (x Config) Set(key string, value string) {
	values := Config.List()
	values[key] = value
}

// List function
func (x Config) List() (values map[string]string) {
	values = make(map[string]string)
	values["Name"] = "jane"
	values["LogFile"] = "/home/matt/jane.log"
	return values
}
