package cmd

import (
	"testing"
)

func TestPageTraversal(t *testing.T) {

	t.Run("Simple one step DFS test", func(t *testing.T) {
		balletLinks := []string{"/wiki/Dance", "/wiki/Flamenco"}
		testWiki := newSpyWikiLinks(balletLinks)
		graph := NewGraph(2, &testWiki)
		graph.depthFirstSearch("Ballet", "Dance")

		if graph.mapPath["Dance"] != "Ballet" {
			t.Errorf("Did not find the expected path")
		}
	})
}
