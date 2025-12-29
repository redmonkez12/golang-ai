package main

import (
	"math"
	"os"
	"path/filepath"
)

func inExplored(needle Point, haystack []Point) bool {
	for _, n := range haystack {
		if n.Row == needle.Row && n.Col == needle.Col {
			return true
		}
	}
	return false
}

func emptyTmp() {
	directory := "./tmp/"
	dir, _ := os.Open(directory)
	filesToDelete, _ := dir.Readdir(0)

	for index := range filesToDelete {
		f := filesToDelete[index]
		fullPath := filepath.Join(directory, f.Name())
		_ = os.Remove(fullPath)
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func euclideanDist(p, goal Point) float64 {
	return math.Sqrt(float64(p.Row-goal.Row)*float64(p.Row-goal.Row) + float64(p.Col-goal.Col)*float64(p.Col-goal.Col))
}
