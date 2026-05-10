package main

func countConnections(groupSize int) int {
	connections := 0
	for i := 2; i <= groupSize; i++ {
		if i == 2 {
			connections += 1
		} else {
			connections += i - 1
		}

	}
	return connections
}
