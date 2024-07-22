package main

import (
	"maps"
	"slices"
	"testing"
)

func TestDjikstra(t *testing.T) {
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

	t.Run("costs for the whole graph", func(t *testing.T) {
		want := make(map[string]int)
		want["start"] = 0
		want["a"] = 4
		want["b"] = 10
		want["c"] = 18
		want["d"] = 15
		want["e"] = 20
		want["fin"] = 24

		got := DjikstraCosts(graph, "start")

		if !maps.Equal(got, want) {
			t.Errorf("\ngot %v,\nwant %v", got, want)
		}
	})

	t.Run("Path for start and finish", func(t *testing.T) {
		want := []string{"start", "b", "d", "e", "fin"}
		got := DjikstraPath(graph, "start", "fin")

		if !slices.Equal(got, want) {
			t.Errorf("\ngot %v,\nwant %v", got, want)
		}
	})

	t.Run("Path for non existent node", func(t *testing.T) {
		want := []string{}
		got := DjikstraPath(graph, "start", "notNode")

		if !slices.Equal(got, want) {
			t.Errorf("\ngot %v,\nwant %v", got, want)
		}
	})
}