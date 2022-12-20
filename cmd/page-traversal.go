package cmd

import "fmt"

// Max pages we will allow the program to visit
// if not controlled traversals could go on and on
const maxPagesVisited = 20

type Graph struct {
	mapPath  map[string]string
	maxPages int
	wiki     WikiLinks
}

func NewGraph(maxPages int, wiki WikiLinks) Graph {
	return Graph{
		mapPath:  make(map[string]string),
		maxPages: maxPages,
		wiki:     wiki,
	}
}

func (g *Graph) depthFirstSearch(start, goal string) (int, bool) {
	// Initialize necessary data structures for DFS
	stack := &Stack{}
	stack.push(start)

	visitedSet := make(map[string]bool)
	visitedSet[start] = true

	index := 0
	var goalReached = false

	for !stack.isEmpty() && index < g.maxPages {
		curr, _ := stack.pop()
		if curr == goal {
			goalReached = true
			break
		}
		neighborLinks := getPageContent(curr, g.wiki)
		for _, link := range neighborLinks {
			if !visitedSet[link] {
				visitedSet[link] = true
				g.mapPath[link] = curr
				stack.push(link)
			}
		}
		index++
	}

	return index, goalReached
}

func (g *Graph) breathFirstSearch(start, goal string) (int, bool) {
	// Initliaze data structures for BFS
	queue := &Queue{}
	queue.enqueue(start)

	visitedSet := make(map[string]bool)
	visitedSet[start] = true

	index := 0
	var goalReached = false

	for !queue.isEmpty() && index < g.maxPages {
		curr, _ := queue.dequeue()
		if curr == goal {
			goalReached = true
			break
		}
		neighborLinks := getPageContent(curr, g.wiki)
		for _, link := range neighborLinks {
			if !visitedSet[link] {
				visitedSet[link] = true
				g.mapPath[link] = curr
				queue.enqueue(link)
			}
		}
		index++
	}

	return index, goalReached

}

func (g *Graph) printPath(start, goal string) int {
	if len(g.mapPath) == 0 {
		return 0
	}

	curr := goal
	count := 0

	for curr != "" {
		fmt.Println(curr)
		curr = g.mapPath[curr]
		count++
	}

	return count
}
