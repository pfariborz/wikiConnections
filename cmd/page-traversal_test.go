package cmd

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"
)

// Mock implementation of getLinks for
// DFS/BFS testing this mock prevents
// unnecessary calls to wikipedia
// and guarantees consistent return values
// for test stability
type mockWikiLinkSearch struct {
	mock.Mock
}

func newMockWikiLink() mockWikiLinkSearch {
	return mockWikiLinkSearch{}
}

func (m *mockWikiLinkSearch) getLinks(link string) []string {
	args := m.Called(link)
	return args.Get(0).([]string)
}

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

	t.Run("DFS and BFS test", func(t *testing.T) {
		// Start and Goal for DFS and BFS
		start := "Spain"
		goal := "Madrid"

		//DFS Graph
		testWikiDfs := newMockWikiLink()
		firstLinks := []string{"/wiki/Argentina", "/wiki/Spanish", "/wiki/Europe"}
		secondLinks := []string{"/wiki/Barcelona", "/wiki/Tapas", "/wiki/Wine"}
		thirdLinks := []string{"/wiki/Soccer", "/wiki/Madrid"}
		testWikiDfs.On("getLinks", mock.Anything).Return(firstLinks).Once()
		testWikiDfs.On("getLinks", mock.Anything).Return(secondLinks).Once()
		testWikiDfs.On("getLinks", mock.Anything).Return(thirdLinks)

		graphDfs := NewGraph(20, &testWikiDfs)
		hopsDfs, goalReachedDfs := graphDfs.depthFirstSearch(start, goal)
		if hopsDfs != 3 {
			t.Errorf("DFS search alogrithm for this output should visit 3 hops on path from Soccer -> Madrid")
		}
		if !goalReachedDfs {
			t.Errorf("Expected to reach the goal, instead DFS reported goal was not reached")
		}

		pathCountDfs := graphDfs.printPath(start, goal)
		if pathCountDfs != 4 {
			t.Errorf("Expected path to be 4 total hops for DFS")
		}

		//BFS Graph
		testWikiBfs := newMockWikiLink()
		testWikiBfs.On("getLinks", mock.Anything).Return(firstLinks).Once()
		testWikiBfs.On("getLinks", mock.Anything).Return(secondLinks).Once()
		testWikiBfs.On("getLinks", mock.Anything).Return(thirdLinks)

		graphBfs := NewGraph(20, &testWikiBfs)
		hopsBfs, goalReachedBfs := graphBfs.breathFirstSearch(start, goal)
		if hopsBfs != 8 {
			t.Errorf("BFS search algorithm for this output should visit 8 hops on path from Soccer -> Madrid")
		}
		if !goalReachedBfs {
			t.Errorf("Expected to reach the goal, instead BFS reported goal was not reached")
		}

		pathCountBfs := graphBfs.printPath(start, goal)
		if pathCountBfs != 3 {
			t.Errorf("Expected path to be 3 total hops for BFS")
		}
	})
}
