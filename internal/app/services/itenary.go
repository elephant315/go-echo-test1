package services

import "errors"

// ReconstructItinerary reconstructs the complete flight itinerary from a slice of ticket tuples
func ReconstructItinerary(tickets [][]string) ([]string, error) {
	if len(tickets) == 0 {
		return nil, errors.New("no tickets provided")
	}

	// Map to find the starting point
	destinationMap := make(map[string]string)
	sourceMap := make(map[string]bool)

	for _, ticket := range tickets {
		src, dst := ticket[0], ticket[1]
		destinationMap[src] = dst
		sourceMap[dst] = true
	}

	// Find the starting airport
	var start string
	for src := range destinationMap {
		if !sourceMap[src] {
			start = src
			break
		}
	}

	if start == "" {
		return nil, errors.New("invalid input")
	}

	// Construct the itinerary
	var itinerary []string
	current := start
	for current != "" {
		itinerary = append(itinerary, current)
		next, exists := destinationMap[current]
		if !exists {
			break
		}
		current = next
	}

	return itinerary, nil
}
