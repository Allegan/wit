package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Config holds the structure of a JSON configuration file which is loaded
type Config struct {
	Protocol string `json:"protocol"`
	Host     string `json:"host"`
	API      string `json:"api"`
}

// exists returns whether the given file or directory exists
// https://stackoverflow.com/questions/10510691/how-to-check-whether-a-file-or-directory-exists
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// ReadFromFile takes a filePath string, reads the content, and unmarshals it
// into an instance of the Config struct
func (config *Config) ReadFromFile(filePath string) (c *Config, err error) {
	// check if file exists
	exist, err := exists(filePath)
	if !exist {
		config.SaveToFile(filePath)
	}

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("Failed to read `%v`: \n\t%w", filePath, err)
	}

	err = json.Unmarshal([]byte(content), config)
	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshall config: \n\t%w", err)
	}

	return config, nil
}

// SaveToFile takes a filePath string, marshals the struct, and saves the
// string to the corresponding file
func (config *Config) SaveToFile(filePath string) (err error) {
	content, err := json.Marshal(config)
	if err != nil {
		return fmt.Errorf("Failed to marshall config: \n\t%w", err)
	}

	err = ioutil.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("Failed to write `%v`: \n\t%w", filePath, err)
	}

	return nil
}
