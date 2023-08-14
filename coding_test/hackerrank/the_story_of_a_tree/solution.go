package the_story_of_a_tree

import (
	"fmt"
)

/**
 * The Story of a Tree
 *
 * https://www.hackerrank.com/challenges/the-story-of-a-tree/problem?isFullScreen=false
 **/

type Node struct {
	index  int32
	weight int32
}

func storyOfATree(n int32, edges [][]int32, k int32, guesses [][]int32) string {
	caseCountArr := make([]int32, n+1)
	edgeArr := make([][]int32, n+1)
	guessMap := make(map[int32]map[int32]bool, 0)
	var result string

	for _, edge := range edges {
		n1, n2 := edge[0], edge[1]

		if edgeArr[n1] == nil {
			edgeArr[n1] = make([]int32, 0)
		}
		if edgeArr[n2] == nil {
			edgeArr[n2] = make([]int32, 0)
		}

		edgeArr[n1], edgeArr[n2] = append(edgeArr[n1], n2), append(edgeArr[n2], n1)
	}

	for _, guess := range guesses {
		from, to := guess[1], guess[0]
		if _, ok := guessMap[from]; !ok {
			guessMap[from] = make(map[int32]bool)
		}

		guessMap[from][to] = false
	}

	queue := make([]Node, 0)
	for _, guess := range guesses {
		from, to := guess[1], guess[0]

		if guessMap[from][to] {
			continue
		}

		guessMap[from][to] = true
		visitedNodeArr := make([]bool, n+1)
		visitedNodeArr[from] = true
		queue = append(queue, Node{index: to, weight: 1})

		for len(queue) != 0 {
			fromNode := queue[0]
			queue = queue[1:]
			caseCountArr[fromNode.index] += fromNode.weight
			visitedNodeArr[fromNode.index] = true

			for _, to := range edgeArr[fromNode.index] {
				if !visitedNodeArr[to] {
					nextNode := Node{index: to, weight: fromNode.weight}

					if isPassed, ok := guessMap[fromNode.index][to]; !isPassed && ok {
						nextNode.weight++
						guessMap[fromNode.index][to] = true
					}

					queue = append(queue, nextNode)
				}
			}
		}
	}

	rootCanCount := int32(0)
	for i := int32(1); i < n+1; i++ {
		if k <= caseCountArr[i] {
			rootCanCount++
		}
	}

	if rootCanCount == n {
		result = "1/1"
	} else if rootCanCount == 0 {
		result = "0/1"
	} else {
		for i := int32(2); i <= rootCanCount && i <= n; {
			if rootCanCount%i == 0 && n%i == 0 {
				rootCanCount /= i
				n /= i
			} else {
				i++
			}
		}

		result = fmt.Sprintf("%d/%d", rootCanCount, n)
	}

	return result
}
