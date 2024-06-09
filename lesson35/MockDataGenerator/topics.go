package mockdatagenerator

import (
	"fmt"
	"leetcode/model"
	"leetcode/storage/postgres"
)

func TopicsGenerator(t *postgres.Topics) error {
	topics := []string{"Array", "String", "Hash Table", "Math", "Sorting", "Greedy",
		"Binary Search", "Divide and Conquer", "Dynamic Programming", "Backtracking",
		"Stack", "Heap (Priority Queue)", "Linked List", "Tree", "Trie", "Graph",
		"Depth-First Search", "Breadth-First Search", "Union Find", "Bit Manipulation",
		"Sliding Window", "Two Pointers", "Intervals", "Monotonic Stack",
		"Ordered Set", "Enumeration", "Matrix", "Geometry", "Simulation",
		"Reservoir Sampling", "Memoization", "Shortest Path", "Topological Sort",
		"Minimum Spanning Tree", "Probability and Statistics", "Game Theory",
		"Combinatorics"}

	for _, topicName := range topics {
		newTopic := model.Topic{
			Name: topicName,
		}
		err := t.CreateTopic(&newTopic)
		if err != nil {
			return err
		}
	}
	fmt.Println("Barcha topiclar qo'shildi")
	return nil
}
