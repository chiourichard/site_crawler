package engine

import (
    "fmt"
    "net/http"
    "golang.org/x/net/html"
    "log"
    "net/url"
    "regexp"
)

var tokens = make(chan struct{}, 20)
var regex string = ""
var list []string

func crawl(seedDomainName string, webUrl string) []string {
    fmt.Println(webUrl)
    tokens <- struct{}{} // acquire a token
    list, err := Extract(webUrl)
    <-tokens // release the token
    if err != nil {
        log.Print(err)
    }
    return list
}

func Extract(seedDomainName string, webUrl string) ([]string, error) {
    resp, err := http.Get(webUrl)
    if err != nil {
        return nil, err
    }
    if resp.StatusCode != http.StatusOK {
    resp.Body.Close()
        return nil, fmt.Errorf("getting %s: %s", webUrl, resp.Status)
    }
    doc, err := html.Parse(resp.Body)
    resp.Body.Close()
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
                link, err := resp.Request.URL.Parse(a.Val)
                if err != nil {
                    continue // ignore bad URLs
                }
                if isSameDomain(seedDomainName,link.String()){
                    links = append(links, link.String())
                }
            }
        }
    }
    forEachNode(doc, visitNode, nil)

    return links, nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
    if pre != nil {
        pre(n)
    }
    for c := n.FirstChild; c != nil; c = c.NextSibling {
        forEachNode(c, pre, post)
    }
    if post != nil {
        post(n)
    }
}

func isValidUrl(toTest string) bool {
    var validHttp = regexp.MustCompile("^(http|https)://")
    var matchHttp = validHttp.MatchString(toTest)

    _, err := url.ParseRequestURI(toTest)
    if err != nil || matchHttp == false {
        return false
    } else {
        return true
    }
}