package main

func countConnections(groupSize int) int {
	connections := 0
	for person := 0; person < groupSize; person++ {
		connections += person
	}
	return connections
}
