package cmd

import (
	"reflect"
	"testing"
)

func TestPageTraversal(t *testing.T) {

	t.Run("Simple one step DFS test", func(t *testing.T) {
		balletLinks := []string{"/wiki/Dance", "/wiki/Flamenco"}
		testWiki := newSpyWikiLinks(balletLinks)
		graph := NewGraph(2, &testWiki)
		graph.depthFirstSearch("Ballet", "Dance")

		expectedMap := map[string]string{"Dance": "Ballet", "Flamenco": "Ballet"}
		if !reflect.DeepEqual(expectedMap, graph.mapPath) {
			t.Errorf("Did not get expected mapPath for Ballet -> Dance")
		}
	})
}
