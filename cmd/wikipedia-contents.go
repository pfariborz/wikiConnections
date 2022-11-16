package cmd

import (
	"log"
	"net/http"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

const wikiLink = "https://en.wikipedia.org/wiki/"
const wikiPrefix = "/wiki/"

// Public facing routine usable by outside packages
// Input: Subject of interest in Wikipedia
// Examples include: Basketball, Starfish, Hawaii
// Output: array of internal wiki links from the subject page
func GetPageLinks(subject string) []string {
	// WikiLink struct will go to wikipedia.org directly
	wiki := &WikiLinkSearch{}
	return getPageContent(subject, wiki)
}

// Input: subject of interest in Wikipedia
//        wikiLinks in order to allow for testing a mock of wikiLinks
//        has been made, see SpyWikiLinks for implementation
func getPageContent(subject string, wikiLinks WikiLinks) []string {
	if len(subject) == 0 || subject == "" {
		return nil
	}
	// Take the subject passed and add it to the wiki link
	link := wikiLink + subject
	urls := getInternalWikiLinks(link, wikiLinks)

	return urls
}

// Filters out unneccessary common internal wiki links to
// ensure search is only looking for real subject connections
// between pages
func getInternalWikiLinks(link string, wikiLinks WikiLinks) []string {
	var internalWikiLinks []string

	// Collect all links from page
	pageLinks := wikiLinks.getLinks(link)
	for _, link := range pageLinks {
		match, _ := regexp.MatchString("^/wiki/.+:", link)
		if match {
			continue
		}

		// Bad data sent to Wikipedia will result in the following generic internal links
		if strings.HasPrefix(link, "/wiki/Main_Page") || strings.HasPrefix(link, "/wiki/Case_sensitivity") {
			continue
		}

		if strings.HasPrefix(link, wikiPrefix) {
			// Remove internal link prefix
			cleanedLink := strings.Replace(link, "/wiki/", "", -1)
			internalWikiLinks = append(internalWikiLinks, cleanedLink)
		}
	}

	return internalWikiLinks
}

type WikiLinks interface {
	getLinks(link string) []string
}

// General WikiLinkSearch struct used for runtime code
// goes directly to wikipedia.org for getLinks implementation
type WikiLinkSearch struct{}

// Collect all links from Wikipedia subject response body and returns
// an array of those internal Wikipedia links
func (w *WikiLinkSearch) getLinks(link string) []string {

	res, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	defer res.Body.Close()

	var links []string
	z := html.NewTokenizer(res.Body)
	for {
		tt := z.Next()

		switch tt {
		case html.ErrorToken:
			//todo: links list shoudn't contain duplicates
			return links
		case html.StartTagToken, html.EndTagToken:
			token := z.Token()
			if "a" == token.Data {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						links = append(links, attr.Val)
					}
				}
			}
		}
	}
}

// Mock struct for unit testing
type SpyWikiLinks struct {
	links []string
}

// Mock links
var spyLinks = [4]string{"/wiki/Rugby_sevens", "/wiki/Test:template", "/wiki/Tennis", "/wiki/Document:Tennis.jpg"}

// Mock implementation of getLinks used for unit testing
// instead of hitting wikipedia.org
// will return back a list of mocked internal links
func (s *SpyWikiLinks) getLinks(link string) []string {
	for _, v := range spyLinks {
		s.links = append(s.links, v)
	}
	return s.links
}
