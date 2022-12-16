package cmd

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"
)

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

	t.Run("Multiple calls for DFS test", func(t *testing.T) {
		//DFS Graph
		testWikiDfs := newMockWikiLink()
		firstLinks := []string{"/wiki/Argentina", "/wiki/Spanish", "/wiki/Europe"}
		secondLinks := []string{"/wiki/Barcelona", "/wiki/Tapas", "/wiki/Wine"}
		thirdLinks := []string{"/wiki/Soccer", "/wiki/Madrid"}
		testWikiDfs.On("getLinks", mock.Anything).Return(firstLinks).Once()
		testWikiDfs.On("getLinks", mock.Anything).Return(secondLinks).Once()
		testWikiDfs.On("getLinks", mock.Anything).Return(thirdLinks)

		graphDfs := NewGraph(20, &testWikiDfs)
		hopsDfs := graphDfs.depthFirstSearch("Spain", "Madrid")
		if hopsDfs != 3 {
			t.Errorf("DFS search alogrithm for this output should return 3 hops to path from Soccer -> Madrid")
		}

		//BFS Graph
		testWikiBfs := newMockWikiLink()
		testWikiBfs.On("getLinks", mock.Anything).Return(firstLinks).Once()
		testWikiBfs.On("getLinks", mock.Anything).Return(secondLinks).Once()
		testWikiBfs.On("getLinks", mock.Anything).Return(thirdLinks)

		graphBfs := NewGraph(20, &testWikiBfs)
		hopsBfs := graphBfs.breathFirstSearch("Spain", "Madrid")
		if hopsBfs != 8 {
			t.Errorf("BFS search algorithm for this output should return 8 hops to path from Soccer -> Madrid")
		}
	})
}
