package engine

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"golang.org/x/net/html"
)

var tokens = make(chan struct{}, 20)
var regex string = ""
var FolderName string = ""
var CrawledUrls []string = nil

// crawl a web page
func Crawl(webUrl string) []string {
	fmt.Println(webUrl)
	CrawledUrls = append(CrawledUrls, webUrl)
	err := DownloadFile(FolderName+"/"+url.PathEscape(webUrl), webUrl)
	if err != nil {
		log.Print(err)
	}
	tokens <- struct{}{} // acquire a token
	webUrlList, err := Extract(webUrl)
	<-tokens // release the token
	if err != nil {
		log.Print(err)
	}
	return webUrlList
}

// Extract same domain links with seed from a web page
func Extract(webUrl string) ([]string, error) {
	response, err := http.Get(webUrl)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		response.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", webUrl, response.Status)
	}
	doc, err := html.Parse(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", webUrl, err)
	}
	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := response.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				if IsSameDomain(link.String()) {
					links = append(links, link.String())
				}
			}
		}
	}
	ForEachNode(doc, visitNode, nil)

	return links, nil
}

// explore nodes in a page
func ForEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ForEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}
