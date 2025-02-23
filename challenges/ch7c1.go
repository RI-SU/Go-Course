package main

func countConnections(groupSize int) int {
	connections := 0
	connectors := groupSize - 1
	for i := 1; i <= groupSize; i++ {
		connections += connectors
		connectors--
	}
	return connections
}
