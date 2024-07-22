package main

import (
	"fmt"
	"math"
	"slices"
)

type Graph map[string]map[string]int

func main() {
	graph := make(Graph)

	graph["start"] = make(map[string]int)
	graph["start"]["a"] = 4
	graph["start"]["b"] = 10

	graph["a"] = make(map[string]int)
	graph["a"]["e"] = 21

	graph["b"] = make(map[string]int)
	graph["b"]["d"] = 5
	graph["b"]["c"] = 8

	graph["c"] = make(map[string]int)
	graph["c"]["e"] = 12

	graph["d"] = make(map[string]int)
	graph["d"]["e"] = 5

	graph["e"] = make(map[string]int)
	graph["e"]["fin"] = 4

	graph["fin"] = make(map[string]int)

	costs := DjikstraCosts(graph, "start")
	for f, s := range costs {
		fmt.Printf("%v : %v\n", f, s)
	}

	path := DjikstraPath(graph, "start", "fin")
	fmt.Printf("path: %v\n", path)
}

func LeastCost(cost map[string]int, marked map[string]bool) (string, bool) {
	var least string
	minCost := math.MaxInt
	allMarked := true

	for f, s := range cost {
		if s < minCost && !marked[f] {
			minCost = s
			least = f
			allMarked = false
		}
	}

	return least, allMarked
}

func Djikstra(graph Graph, start string) (map[string]int, map[string]string) {
	cost := make(map[string]int)
	parent := make(map[string]string)
	marked := make(map[string]bool)

	for f := range graph {
		cost[f] = math.MaxInt
	}
	cost[start] = 0

	for {
		node, allMarked := LeastCost(cost, marked)
		if allMarked {
			break
		}

		marked[node] = true

		for f, s := range graph[node] {
			if cost[node] + s < cost[f] {
				cost[f] = cost[node] + s
				parent[f] = node
			}
		}
	}

	return cost, parent
}

func DjikstraCosts(graph Graph, start string) map[string]int {
	costs,_ := Djikstra(graph, start)

	return costs
}

func DjikstraPath(graph Graph, start, goal string) []string {
	if graph[start] == nil || graph[goal] == nil {
		return []string{}
	}
	path := []string{}

	_, parent := Djikstra(graph, start)

	node := goal
	for {
		path = append(path, node)
		if node == start {
			break
		} else {
			node = parent[node]
		}
	}
	slices.Reverse(path)

	return path
}