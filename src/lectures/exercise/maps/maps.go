//--Summary:
//  Create a program to display server status. The server statuses are
//  defined as constants, and the servers are represented as strings
//  in the `servers` slice.
//
//--Requirements:
//* Create a function to print server status displaying:
//  - number of servers
//  - number of servers for each status (Online, Offline, Maintenance, Retired)
//* Create a map using the server names as the key and the server status
//  as the value

package main

import "fmt"

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

func printStatus(status map[string]int) {
	amount := len(status)
	fmt.Println(amount, "servers")
	statuses := make(map[string]int)
	for _, element := range status {
		switch element {
		case Online:
			statuses["Online"]++
		case Offline:
			statuses["Offline"]++
		case Maintenance:
			statuses["Maintenance"]++
		case Retired:
			statuses["Retired"]++
		}
	}
	for i, element := range statuses {
		fmt.Println(element, i, "servers")
	}

}

func main() {
	//* Set all of the server statuses to `Online` when creating the map
	//* After creating the map, perform the following actions:
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}
	status := make(map[string]int)
	for _, element := range servers {
		status[element] = Online
	}
	//  - call display server info function
	printStatus(status)
	//  - change server status of `darkstar` to `Retired`
	status["darkstar"] = Retired
	//  - change server status of `aiur` to `Offline`
	status["aiur"] = Offline
	//  - call display server info function
	printStatus(status)
	//  - change server status of all servers to `Maintenance`
	for i, _ := range status {
		status[i] = Maintenance
	}
	printStatus(status)
	//  - call display server info function

}
