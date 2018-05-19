package engine

import (
    "fmt"
    "net/http"
    "golang.org/x/net/html"
    "log"
    "strings"
)

var tokens = make(chan struct{}, 20)
var regex string = ""
var FolderName string = ""

func Crawl(webUrl string) []string {
    fmt.Println(webUrl)
   // err := DownloadFile(FolderName + "/" + strings.Replace(webUrl, "/", "_", -1), webUrl)
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

func Extract(webUrl string) ([]string, error) {
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
                if IsSameDomain(link.String()){
                    links = append(links, link.String())
                }
            }
        }
    }
    ForEachNode(doc, visitNode, nil)

    return links, nil
}

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
