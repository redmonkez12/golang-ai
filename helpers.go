package main

func isExplored(needle Point, haystack []Point) bool {
	for _, n := range haystack {
		if n.Row == needle.Row && n.Col == needle.Col {
			return true
		}
	}
	return false
}
