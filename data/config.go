package data

import (
	"errors"
)

//Config struct
type Config struct {
	KeyValues map[string]string
}

// Get function
func (x Config) Get(key string) (string, error) {
	if val, ok := x.KeyValues[key]; ok {
		return val, nil
	}

	err := errors.New("Key does not exist.")
	return "", err
}

// Set function
func (x Config) Set(key string, value string) {
	x.KeyValues[key] = value
}

// List function
func (x Config) List() map[string]string {
	return x.KeyValues
}
