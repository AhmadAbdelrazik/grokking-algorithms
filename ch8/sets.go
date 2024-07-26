package main

import "maps"

func Union(set1, set2 set) set {
	resultSet := make(set)

	for val := range set1 {
		resultSet[val] = true
	}
	for val := range set2 {
		resultSet[val] = true
	}

	return resultSet
}

func Intersect(set1, set2 set) set {
	resultSet := make(set)

	if len(set1) > len(set2) {
		set1, set2 = set2, set1
	}

	for val := range set1 {
		if set2[val] {
			resultSet[val] = true
		}
	}

	return resultSet
}

func Difference(set1, set2 set) set {
	resultSet := make(set)

	for val := range set1 {
		if !set2[val] {
			resultSet[val] = true
		}
	}

	return resultSet
}

func Equal(set1, set2 set) bool {
	return maps.Equal(set1, set2)
}