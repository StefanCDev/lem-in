package lemin

// func LinkedList() {
// 	for _, v := range Node.Children {

// 		// if v.CurrentLevel != 0 {
// 		// // 	Node.Parents = append(Node.Parents, v)
// 		// // }

// 		//add the parent node to the child
// 		//only if it doesn't already exist

// 		if !inArray(v.Parents, Node) {
// 			v.Parents = append(v.Parents, Node)
// 			for index, item := range v.Children {
// 				if item == Node && index != len(v.Children)-1 {
// 					v.Children = append(v.Children[:index], v.Children[index+1:]...)
// 				} else if item == Node && index == len(v.Children)-1 {
// 					v.Children = v.Children[:index]
// 				}
// 			}
// 		} else {
// 			continue
// 		}

// 		//remove the Parent node from the LinkedList

// 		//Result.Head = v

// 	}
// }
// func AdjustMap(MyMap *Map) {

// 	//Starting room is the first entry in the LinkedList
// 	for key := range *MyMap {
// 		if key.Start {
// 			//the starting point will always be Tail
// 			//Head will move across the map
// 			Result.Tail, Result.Head = key, key
// 		}
// 		if key.End {
// 			key.Parents = key.Children
// 			key.Children = nil
// 		}
// 	}

// }

// // Should implement the level system (increment it by 1 for each edge between it and start)
// func LinkedList(Node *Vertice) {
// 	Node.Visited = true
// 	for i, v := range Node.Children {
// 		if v.Visited {
// 			Node.Parents = append(Node.Parents, v)
// 			Node.Children[i] = nil
// 		}
// 	}
// 	LinkedList(N)
// }

// // create a map[string]int
// // string represent the ants as keys
// // int to keep track of their pathing until end is reached
// PrintIndex := make(map[string]int)

// //Add all ants to the map
// for index, Q := range QueuedAnts {
// 	for _, Ant := range Q {
// 		PrintIndex[Ant] = len(MaxFlow[index]) - 2
// 	}
// }

// //each loop == number of turn
// //Print Turn 1:
// //Print Turn 2:
// //...
// //...
// //Print Turn N:

// var MovingAnts []Ants

// 	fmt.Println("turn", i, ":")
// 	//i represents the number of turns
// 	// the relationship it shares with ants
// 	// is the number of Ants to initialize from the Q

// 	//cap Add to the len of the ants' respective path
// 	var Add int = 0
// // 	for index, Q := range QueuedAnts {
// 		if Add < len(Q) {
// 			MovingAnts = append(MovingAnts, Q[Add])

// 		}
// 		if index == len(QueuedAnts)-1 {
// 			for _, ant := range MovingAnts {

// 				RelativePosition := MaxFlow[index]
// 				value, exists := PrintIndex[ant]

// 				if exists {
// 					Result += fmt.Sprintf("%s-%s ", ant, RelativePosition[len(RelativePosition)-value].Name)
// 					PrintIndex[ant]--
// 					if value < 1 {
// 						delete(PrintIndex, ant)
// 					}
// 				}

// 			}
// 			Result += "\n"
// 		}
// 	}
// 	Add++

// }

// // initialize the ants from the queue one by one,
// // ([]string and delete elements that reached the end)
// // by incrementing the index from QueuedAnts

// for _, Path := range MaxFlow {
// 	//Recursive for each ant L inside QueuedAnts until all reach Node.End?
// }
// return Result
