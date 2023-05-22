package graph

func IsBipartite(graph [][]int) bool {
	size := len(graph)
	queue := make([]int, 0)
	color := make([]int, size)
	currentColor := 1
	for index, _ := range graph {
		if color[index] == 0 {
			queue = append(queue, index)
		}
		color[index] = currentColor
		for len(queue) > 0 {

			currentIndex := dequeue(&queue)
			if color[currentIndex] == 1 {
				currentColor = 2
			} else {
				currentColor = 1
			}
			for _, value2 := range graph[currentIndex] {
				if color[value2] == 0 {
					color[value2] = currentColor
					queue = append(queue, value2)
				} else if color[value2] != currentColor {
					return false
				}
			}
		}
	}

	return true
}

func dequeue(queue *[]int) int {
	value := (*queue)[0]
	*queue = (*queue)[1:]
	return value
}
