package cmd

// Max pages we will allow the program to visit
// if not controlled traversals could go on and on
const maxPagesVisited = 50

type Graph struct {
	mapPath  map[string]string
	maxPages int64
	wiki     WikiLinks
}

func NewGraph(maxPages int64, wiki WikiLinks) Graph {
	return Graph{
		mapPath:  make(map[string]string),
		maxPages: maxPages,
		wiki:     wiki,
	}
}

func (g *Graph) depthFirstSearch(start string, goal string) {
	// Initialize necessary data structures for DFS
	stack := &Stack{}
	stack.push(start)

	visitedSet := make(map[string]bool)
	visitedSet[start] = true

	index := 0

	for !stack.isEmpty() && index < maxPagesVisited {
		curr, _ := stack.pop()
		if curr == goal {
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
}
