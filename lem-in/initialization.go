package lemin

import (
	"log"
	"reflect"
	"sort"
)

var AllPaths [][]Vertex

type Vertex struct {
	Name  string
	Start bool
	End   bool
	Links []*Vertex
}

type Flow struct {
	PathComb   [][]Vertex
	TotalTurns int
}

// STEP 1: Get all possible individual paths
func RecursivePathFinder(Node *Vertex, route []Vertex) {
	if Node.End {
		route = append(route, *Node)
		sepRoute := make([]Vertex, len(route))
		copy(sepRoute, route)
		AllPaths = append(AllPaths, sepRoute)
		return
	}

	// Check whether we have added current vertex to the route
	if inArray(route, *Node) {
		return
	}

	// If we haven't added current vertex to the route - add it now
	route = append(route, *Node)

	for _, v := range Node.Links {
		RecursivePathFinder(v, route)
	}
}

// STEP 2: find all combinations of unique paths
func CombinePaths(AllPossiblePaths [][]Vertex) []Flow {
	//THE UGLIEST FUNCTION I HAVE EVER WRITTEN
	if len(AllPossiblePaths) == 0 {
		log.Fatal("no path found from start to end")
	}

	sort.Slice(AllPossiblePaths, func(i, j int) bool { return len(AllPossiblePaths[j]) > len(AllPossiblePaths[i]) })

	Result := make([]Flow, 0)
	CombPaths := make([][]Vertex, 0)
	var Breaker bool

	for _, P1 := range AllPossiblePaths {
		//compare 1 Path with all other Paths
		CombPaths = append(CombPaths, P1)
		for _, P2 := range AllPossiblePaths {
			//compare the Paths node by node.
			for i, P := range CombPaths {
				if i == len(CombPaths)-1 && !Breaker {
					CombPaths = append(CombPaths, P2)
				}
				if !Breaker {
					for _, v := range P[1 : len(P)-1] {
						if inArray(P2[1:len(P2)-1], v) {
							Breaker = true
							break
						}
					}
				}
			}
			Breaker = false
		}

		Result = append(Result, Flow{PathComb: CombPaths})

		CombPaths = nil
	}
	return Result
}

// func FindAllPaths(graph map[*Vertex]struct{}, source, sink *Vertex) [][]Vertex {
// 	visited := make(map[*Vertex]bool)
// 	paths := [][]Vertex{}
// 	currentPath := []Vertex{}

// 	recursivePathFinder(source, sink, visited, currentPath, &paths)

// 	return paths
// }

// func recursivePathFinder(current *Vertex, sink *Vertex, visited map[*Vertex]bool, currentPath []Vertex, paths *[][]Vertex) {
// 	visited[current] = true
// 	currentPath = append(currentPath, *current)

// 	if current == sink {
// 		// Append a copy of the current path to the list of paths
// 		copiedPath := make([]Vertex, len(currentPath))
// 		copy(copiedPath, currentPath)
// 		*paths = append(*paths, copiedPath)
// 	} else {
// 		for _, neighbor := range current.Links {
// 			if !visited[neighbor] {
// 				recursivePathFinder(neighbor, sink, visited, currentPath, paths)
// 			}
// 		}
// 	}

// 	// Backtrack
// 	// currentPath = currentPath[:len(currentPath)-1]
// 	visited[current] = false
// }

// STEP 3: Find the Maximum Flow
func (F *Flow) FindTotalTurns(TotalAnts int) {

	TotalFlow := len(F.PathComb)
	QueuedAnts := make([][]string, TotalFlow)

	sort.Slice(F.PathComb, func(i, j int) bool { return len(F.PathComb[j]) > len(F.PathComb[i]) })

	//simulate the Edmonds-Karp algorithm for each combination of paths found
	for i := 1; i <= TotalAnts; i++ {
		for j := 0; j < TotalFlow; j++ {
			if j < TotalFlow-1 {
				PathSize1 := len(F.PathComb[j]) + len(QueuedAnts[j])
				PathSize2 := len(F.PathComb[j+1]) + len(QueuedAnts[j+1])
				if PathSize1 <= PathSize2 {
					QueuedAnts[j] = append(QueuedAnts[j], "ant")
					break
				}
			} else if j == TotalFlow-1 {
				QueuedAnts[j] = append(QueuedAnts[j], "ant")
			}

		}
	}

	var temp, index int

	for i, v := range QueuedAnts {
		if len(v) > temp {
			temp = len(v)
			index = i
		}
	}

	F.TotalTurns = len(QueuedAnts[index]) + len(F.PathComb[index]) - 2

}

func inArray(s []Vertex, vp Vertex) (result bool) {
	for _, v := range s {
		if reflect.DeepEqual(v, vp) {
			result = true
			return
		}
	}
	return
}
