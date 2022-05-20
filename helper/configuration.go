package helper

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Shell struct {
	Name string // set to default if it should be the default shell
	Command   string
	Arguments []string // if argument == "$DIR$" it will be replaced by the destination directory
}

type Configuration struct {
	Waypoints []Waypoint
	Shells    []Shell
}

var homeDir, _ = os.UserHomeDir()
var configurationPath = filepath.Join(homeDir, ".easycd.json")

func GetConfiguration() Configuration {
	var currentConfig Configuration

	if _, err := os.Stat(configurationPath); os.IsNotExist(err) {
		os.Create(configurationPath)
		os.WriteFile(configurationPath, []byte("[]"), 0644)
	}

	var config, err = os.ReadFile(configurationPath)
	CheckError(err)

	err = json.Unmarshal([]byte(config), &currentConfig)
	CheckError(err)

	return currentConfig
}

func SaveConfiguration(config Configuration) {
	var marshalledConfig, err = json.Marshal(config)

	CheckError(err)

	if _, err := os.Stat(configurationPath); os.IsNotExist(err) {
		os.Create(configurationPath)
	}

	os.WriteFile(configurationPath, []byte(string(marshalledConfig)), 0644)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}