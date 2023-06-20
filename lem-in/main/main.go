package main

import (
	"fmt"
	"lemin"
	"log"
	"os"
	"sort"
	"strings"
)

var Route []lemin.Vertex
var filename string = os.Args[1]
var NumberAnts int

func main() {
	// Read data from file
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err.Error())
	}
	// Split data on newline character
	content := strings.Split(string(data), "\n")

	// validate first line is number of ants
	NumberAnts = lemin.ValidateAnts(content)

	// Store each room (vertex) in a map
	// and get the index at which the links portion starts
	MyMap, index := lemin.ValidateRooms(content)

	// Make sure there are no duplicate rooms
	lemin.CheckDuplicateNames(&MyMap)

	// Ensure links are valid (ie rooms don't link to themselves, no duplicate links)
	MyMap = lemin.ValidateLinks(content[index:], &MyMap)

	//start recursive path finder function when you find the starting room
	for k := range MyMap {
		if k.Start {
			lemin.RecursivePathFinder(k, Route)
		}
	}

	AllPossibleCombs := lemin.CombinePaths(lemin.AllPaths)

	for i := range AllPossibleCombs {
		AllPossibleCombs[i].FindTotalTurns(NumberAnts)
	}

	//the maxflow is the combination of paths which requires the least amount of turn to output N amount of Ants
	sort.Slice(AllPossibleCombs, func(i, j int) bool { return AllPossibleCombs[i].TotalTurns < AllPossibleCombs[j].TotalTurns })
	MaxFlow := AllPossibleCombs[0].PathComb

	// Queue each ant
	QueuedAnts := lemin.QueueThem(NumberAnts, MaxFlow)
	// Print out result
	fmt.Println(string(data) + "\n")

	lemin.PrintResult(QueuedAnts, MaxFlow, NumberAnts)
}
