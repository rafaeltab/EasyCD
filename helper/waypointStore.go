package helper

import (
	"encoding/json"
	"fmt"
	"os"
)

type Waypoint struct {
	Name string
	Path string
}

var store = `[
	{
		"Name": "home",
		"Path": "/home/user"
	}
]`

func AddWaypoint(name string, directory string) {
	fmt.Println("Creating waypoint", name, "in", directory)

	var currentStore []Waypoint = GetWaypoints()

	currentStore = append(currentStore, Waypoint{
		Name: name,
		Path: directory,
	})	

	saveStore(currentStore)

	fmt.Println(currentStore)
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

	fmt.Println("Waypoint", name, "not found")
	os.Exit(1)
	return Waypoint{}
}

func getStore() []Waypoint {
	var currentStore []Waypoint

	var err = json.Unmarshal([]byte(store), &currentStore)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return currentStore
}

func saveStore(newStore []Waypoint) {
	var marshalledStore, err = json.Marshal(store)

	if(err != nil) {
		fmt.Println(err)
		os.Exit(1)
	}

	store = string(marshalledStore)
}
