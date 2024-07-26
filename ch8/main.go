package main

import "fmt"

type set map[string]bool

func main() {
	state := ToSet([]string{"mt", "wa", "or", "id", "nv", "ut", "ca", "az"})

	stations := make(map[string]set)
	stations["kone"] = ToSet([]string{"id", "nv", "ut"})
	stations["ktwo"] = ToSet([]string{"id", "wa", "mt"})
	stations["kthree"] = ToSet([]string{"or", "nv", "ca"})
	stations["kfour"] = ToSet([]string{"nv", "ut"})
	stations["kfive"] = ToSet([]string{"ca", "az"})

	selectedStations := GetStationsGreedy(state, stations)

	for station := range selectedStations {
		fmt.Printf("station: %v\n", station)
	}

}

func ToSet(arr []string) set {
	set := make(set)
	for _, item := range arr {
		set[item] = true
	}
	return set
}

func GetStationsGreedy(states set, stations map[string]set) set {
	coveredStates := make(set)
	selectedStations := make(set)

	for len(coveredStates) < len(states) {
		highestStation := ""
		newlyCoveredStates := 0

		for station, covers := range stations {
			newStates := make(set)
			for state := range covers {
				if !coveredStates[state] {
					newStates[state] = true
				}
			}

			if len(newStates) > newlyCoveredStates {
				highestStation = station
				newlyCoveredStates = len(newStates)
			}
		}

		if highestStation == "" {
			break
		}

		selectedStations[highestStation] = true
		for states := range stations[highestStation] {
			coveredStates[states] = true
		}
		delete(stations, highestStation)
	}

	return selectedStations
}

// Algorithm

// 1. get the station with the highest uncovered states number
// 2. add the station to the selected stations.
// 3. Mark the states as covered.
// 4. Repeat until the covered states = total states
