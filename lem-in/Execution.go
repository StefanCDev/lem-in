package lemin

import (
	"fmt"
	"sort"
	"strconv"
)

type Ants struct {
	Name  string
	Path  *[]Vertex
	Index int
}

func QueueThem(NumAnts int, MaxFlow [][]Vertex) [][]string {

	//Sort them from shortest to longest
	sort.Slice(MaxFlow, func(i, j int) bool { return len(MaxFlow[j]) > len(MaxFlow[i]) })

	//start queuing them using edmonds-karp
	QueuedAnts := make([][]string, len(MaxFlow))

	//here, we are adding all ants to the only path we have
	//hence why len(MaxFlow) would be 1
	if len(MaxFlow) == 1 {
		for i := 1; i <= NumAnts; i++ {
			AntName := "L"
			QueuedAnts[0] = append(QueuedAnts[0], AntName)
		}
	} else {
		for i := 1; i <= NumAnts; i++ {
			AntName := "L"
			//after adding an ant to the queue
			//we need to decide which path does it
			//correspond to
			for j := 0; j < len(MaxFlow); j++ {
				if j < len(MaxFlow)-1 {
					PathSize1 := len(MaxFlow[j]) + len(QueuedAnts[j])
					PathSize2 := len(MaxFlow[j+1]) + len(QueuedAnts[j+1])
					if PathSize1 <= PathSize2 {
						QueuedAnts[j] = append(QueuedAnts[j], AntName)
						break
					}
				} else if j == len(MaxFlow)-1 {
					QueuedAnts[j] = append(QueuedAnts[j], AntName)
				}

			}
		}

	}

	//Name the ants properly
	counter := 1
	PathLengthCount := 0
	for counter <= NumAnts {
		for _, v := range QueuedAnts {
			if len(v)-1 < PathLengthCount {
				continue
			}
			v[PathLengthCount] += strconv.Itoa(counter)
			counter++
		}
		PathLengthCount++
	}
	return QueuedAnts

}

func PrintResult(QueuedAnts [][]string, MaxFlow [][]Vertex, NumAnts int) {

	var ants []Ants

	var queueCount = len(QueuedAnts)
	var completedQueueCount int = 0

	for i := 0; NumAnts > 0; i++ {
		for j, v := range QueuedAnts {
			if i > len(v)-1 {
				completedQueueCount++
				if completedQueueCount >= queueCount {
					break
				}
			} else {
				ants = append(ants, Ants{Name: v[i], Path: &MaxFlow[j], Index: 1})
			}
		}

		for i, ant := range ants {
			if ant.Index < len(*ant.Path) {
				vertex := *ant.Path
				fmt.Printf("%s-%s ", ant.Name, vertex[ant.Index].Name)
				ant.Index++
				if ant.Index >= len(vertex) {
					NumAnts--
				}
				ants[i] = ant
			}
		}
		fmt.Println()
	}

}
