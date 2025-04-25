package main

func main() {

}

// BFS will find the shortest path (fewest segments)
// Dijkstra's algo will find fastest path with weighted segments
// For example two different segments could each have a different distance
// or time to travel the path, thus making one slower than the other. The path
// with the fewest segments could take the longest and so BFS would not work.
// Dijkstra's algo does not work for negative weight segments. Need Bellman-Ford
// to handle negative weights
