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

var spyLinks = [4]string{"/wiki/Rugby_sevens", "/wiki/Test:template", "/wiki/Tennis", "/wiki/Document:Tennis.jpg"}

type WikiLinks interface {
	getLinks(link string) []string
}

type WikiLinkSearch struct{}

type SpyWikiLinks struct {
	links []string
}

func (s *SpyWikiLinks) getLinks(link string) []string {
	for _, v := range spyLinks {
		s.links = append(s.links, v)
	}
	return s.links
}

func GetPageLinks(subject string) []string {
	wiki := &WikiLinkSearch{}
	return getPageContent(subject, wiki)
}

func getPageContent(subject string, wikiLinks WikiLinks) []string {
	if len(subject) == 0 || subject == "" {
		return nil
	}
	// Take the subject passed and add it to the wiki link
	link := wikiLink + subject
	urls := getInternalWikiLinks(link, wikiLinks)

	return urls
}

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

// Collect all links from response body and return it as an array of strings
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
