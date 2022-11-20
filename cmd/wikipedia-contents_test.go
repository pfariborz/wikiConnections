package cmd

import (
	"reflect"
	"testing"
)

var expectedLinks = []string{"Rugby_sevens", "Tennis"}

func TestWikipediaContents(t *testing.T) {

	t.Run("Happy path verify expected internal links returned", func(t *testing.T) {
		testWord := "Basketball"

		testWiki := newSpyWikiLinks(spyLinks)
		returnedLinks := getPageContent(testWord, &testWiki)
		if !reflect.DeepEqual(returnedLinks, expectedLinks) {
			t.Errorf("Returned internal links do not match expected")
		}
	})

	t.Run("Sending bad data returns nothing", func(t *testing.T) {
		testWord := "blah*$$$qepoijfeqfe"
		wiki := &WikiLinkSearch{}
		returnedLinks := getPageContent(testWord, wiki)
		if len(returnedLinks) > 0 {
			t.Errorf("Expected no internal links returned for bad data")
		}
	})

	t.Run("Sending null data returns nothing", func(t *testing.T) {
		wiki := &WikiLinkSearch{}
		returnedLinks := getPageContent("", wiki)
		if len(returnedLinks) > 0 {
			t.Errorf("Expected empty return for subject of nil passed")
		}
	})

}
