package main

import (
	"fmt"
)

type Event struct {
	Start int
	End   int
}

func findOverlappingEvents(events []Event) (pairs [][]Event) {
	for i := 0; i < len(events); i++ {
		for j := i + 1; j < len(events); j++ {
			if events[i].Start < events[j].End && events[j].Start < events[i].End {
				pairs = append(pairs, []Event{events[i], events[j]})
			}
		}
	}
	return pairs
}

func main() {
	var events []Event

	fmt.Println("Enter the events in the format (start end), type 'exit' to end:")
	for {
		var start, end int
		fmt.Print("Event ", len(events)+1, ": ")
		_, err := fmt.Scan(&start, &end)

		if err != nil {
			if err.Error() == "expected integer" {
				break
			}
			fmt.Println("Error reading input:", err)
			return
		}

		event := Event{Start: start, End: end}
		events = append(events, event)
	}

	overlappingEvents := findOverlappingEvents(events)

	if len(overlappingEvents) == 0 {
		fmt.Println("No overlapping events")
		return
	}

	text := "the event which starts at %d and ends at %d is overlapping the event which starts at %d and ends at %d"
	fmt.Println("Overlapping events:")
	for _, pair := range overlappingEvents {
		fmt.Println(fmt.Sprintf(text, pair[1].Start, pair[1].End, pair[0].Start, pair[0].End))
	}
}
