package amazon

import "fmt"

func CountGroups(related []string) int32 {
	// Write your code here
	fmt.Println(len(related[0]))
	var count int32
	isAdded := make(map[int]bool)
	isVisited := make(map[int]bool)
	for i := 0; i < len(related); i++ {
		if _, ok := isVisited[i]; ok {
			continue
		}
		isVisited[i] = true
		for j := 0; j < len(related[i]); j++ {
			if _, ok := isAdded[i]; !ok {
				isAdded[i] = true
				count++
			}

			if i == j {
				continue
			}
			if related[i][j] == '1' {
				isAdded[j] = true
				isVisited = helper(j, related, isVisited)
			}

		}
	}
	return count
}

func helper(i int, related []string, visited map[int]bool) map[int]bool {
	if _, ok := visited[i]; ok {
		return visited
	}
	visited[i] = true
	for j := 0; j < len(related[i]); j++ {
		if related[i][j] == '1' {
			visited = helper(j, related, visited)
		}
	}
	return visited
}
