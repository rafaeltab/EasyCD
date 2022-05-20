package helper

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Waypoint struct {
	Name string
	Path string
}

func AddWaypoint(name string, directory string) {
	fmt.Println("Creating waypoint", name, "in", directory)

	if(GetWaypoint(name).Name != "") {
		fmt.Println("Waypoint with name", name, "already exists")
		os.Exit(1)
	}

	var currentStore []Waypoint = GetWaypoints()

	currentStore = append(currentStore, Waypoint{
		Name: name,
		Path: directory,
	})	

	saveStore(currentStore)
}

func GetWaypoints() []Waypoint {
	return getStore()
}

func GetWaypoint(name string) Waypoint {
	var currentStore []Waypoint = GetWaypoints()
	for i := 0; i < len(currentStore); i++ {
		if currentStore[i].Name == name {
			return currentStore[i]
		}
	}
	return Waypoint{}
}

var homeDir, _ = os.UserHomeDir()
var storeDir = filepath.Join(homeDir, ".easycd.json")

func getStore() []Waypoint {
	var currentStore []Waypoint

	if _, err := os.Stat(storeDir); os.IsNotExist(err) {
		os.Create(storeDir)
		os.WriteFile(storeDir, []byte("[]"), 0644)
	}

	var store, err = os.ReadFile(storeDir)
	checkError(err)

	err = json.Unmarshal([]byte(store), &currentStore)
	checkError(err)

	return currentStore
}

func saveStore(newStore []Waypoint) {
	var marshalledStore, err = json.Marshal(newStore)

	checkError(err)

	if _, err := os.Stat(storeDir); os.IsNotExist(err) {
		os.Create(storeDir)
	}

	os.WriteFile(storeDir, []byte(string(marshalledStore)), 0644)
}

func checkError(err error){
	if err != nil {
		panic(err)
	}
}