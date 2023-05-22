package main

import (
	"fmt"
	"go_study/main/graph"
)

func main() {
	result := graph.IsBipartite([][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}})

	fmt.Print(result)

}
