package helper

import (
	"fmt"
	"os"
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

func getStore() []Waypoint {
	return GetConfiguration().Waypoints;
}

func saveStore(newStore []Waypoint) {
	var config = GetConfiguration();
	config.Waypoints = newStore;

	SaveConfiguration(config);
}