package main

import (
	"container/list"
	"fmt"
)

func main() {
	routes := map[string][]string{
		"home":       {"harvey", "groff", "silt"},
		"harvey":     {"harvey_gap", "silt_mesa"},
		"groff":      {"harvey_gap"},
		"silt":       {"antlers"},
		"silt_mesa":  {"rifle"},
		"antlers":    {"rifle"},
		"harvey_gap": {"rifle_gap"},
		"rifle":      {"rifle_gap"},
		"rifle_gap":  {},
	}

	result := bfs_graph_search(routes, "home", "rifle_gap")
	fmt.Println(result)
}

type Place struct {
	name     string
	distance int
}

func bfs_graph_search(graph map[string][]string, start string, dest string) (distance int) {
	queue := list.New()
	for _, item := range graph[start] {
		queue.PushBack(Place{item, 1})
	}
	visited := map[string]bool{}

	for queue.Len() > 0 {
		placeEl := queue.Front()
		queue.Remove(placeEl)

		place := placeEl.Value.(Place)
		_, exists := visited[place.name]
		if exists {
			continue
		}

		fmt.Println(place)
		if place.name == dest {
			return place.distance
		} else {
			for _, item := range graph[place.name] {
				queue.PushBack(Place{item, place.distance + 1})
			}
			visited[place.name] = true
		}
	}

	return -1
}
