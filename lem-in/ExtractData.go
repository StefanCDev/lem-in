package lemin

import (
	"log"
	"strconv"
	"strings"
)

var Start, End string = "##start", "##end"
var Scounter, Ecounter int

type Map map[*Vertex]struct{}

var MyMap Map

// Check the first line is a valid number (int and > 0)
func ValidateAnts(content []string) int {
	NumLine1, err := strconv.Atoi(content[0])
	if err != nil || NumLine1 <= 0 {
		log.Fatal("invalid format of inputted ants at line 1")
	}
	_, err = strconv.Atoi(content[1])
	if err == nil {
		log.Fatal("Invalid format of inputted ants at line 2")
	}
	return NumLine1
}

// Returns a map and index integer
func ValidateRooms(data []string) (Map, int) {
	// Create a map of vertices
	MyMap = make(Map)
	var index int
	for i := 1; i < len(data); i++ {
		//reinvent it
		// if room detected
		if strings.Contains(data[i], "-") && len(strings.Fields(data[i])) == 1 {
			index = i
			break
		} else if data[i] == Start || data[i] == End {
			// Check if the next line is a valid room
			if len(strings.Fields(data[i+1])) != 3 || strings.Fields(data[i+1])[0][0] == '#' || strings.Fields(data[i+1])[0][0] == 'L' {
				log.Fatalf("Wrong format of inputted room: %s", data[i+1])
			}
			Room := Vertex{Name: strings.Fields(data[i+1])[0]}
			if data[i] == Start {
				Room.Start = true
				// Check later if only one start
				Scounter++
			} else {
				Room.End = true
				// Check later if only one end
				Ecounter++
			}
			MyMap[&Room] = struct{}{}
			i++
		} else if strings.Fields(data[i])[0][0] == '#' || strings.Fields(data[i])[0][0] == 'L' || len(strings.Fields(data[i])) != 3 {
			log.Fatalf("Wrong format of inputted room: %s", (data[i]))
		} else if len(strings.Fields(data[i])) == 3 {
			MyMap[&Vertex{Name: strings.Fields(data[i])[0]}] = struct{}{}

		}

		// if line == Start or End

		//if 1st line of links block detected

	}

	// If there is more or less than one start/end, throw error
	if Scounter != 1 || Ecounter != 1 {
		log.Fatal("invalid number of start or end indicators")
	}

	return MyMap, index
}

func ValidateLinks(data []string, MyMap *Map) Map {

	// Check if all the called rooms exist
	allLinks := make(map[string][]string)
	links := make(map[string]struct{})

	// Check all links are unique
	for _, item := range data {
		// Search if room links to itself
		temp := strings.Split(item, "-")
		if temp[0] == temp[1] {
			log.Fatal("Room linking to itself")
		}

		// Search for duplicates
		_, ok := allLinks[temp[0]]
		if ok {
			for _, v := range allLinks[temp[0]] {
				if v == temp[1] {
					log.Fatal("Found duplicate link")
				}
			}
		}

		_, ok = allLinks[temp[1]]
		if ok {
			for _, v := range allLinks[temp[1]] {
				if v == temp[0] {
					log.Fatal("Found duplicate link")
				}
			}
		}

		links[temp[0]] = struct{}{}
		links[temp[1]] = struct{}{}
		allLinks[temp[0]] = append(allLinks[temp[0]], temp[1])
	}

	// Check if all the rooms from the links block
	// match the rooms collected from the rooms block

	// Can't be any isolated rooms
	if len(*MyMap) != len(links) {
		log.Fatal("Non-existent or missing room(s) detected within the block of links")
	}

	for k := range *MyMap {
		if _, ok := links[k.Name]; !ok {
			log.Fatalf("Room %s not found", k.Name)
		}
		items := allLinks[k.Name]
		LinksBinder(k, items, MyMap)
	}
	return *MyMap
}

func LinksBinder(Key *Vertex, items []string, MyMap *Map) {

	for _, v := range items {
		for k := range *MyMap {
			if k.Name == v {
				Key.Links = append(Key.Links, k)
				k.Links = append(k.Links, Key)
				break
			}
		}
	}
}

func CheckDuplicateNames(MyMap *Map) {
	NameHolder := make(map[string]bool)

	for k := range *MyMap {
		_, exists := NameHolder[k.Name]
		if exists {
			log.Fatal("duplicate room found")
		} else {
			NameHolder[k.Name] = true
		}
	}

}
