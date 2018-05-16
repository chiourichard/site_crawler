package main  
  
import (
    "os"
    "github.com/chiourichard/golang_site_crawler/engine"
) 
  
var seedDomainName string = ""

func main() {
    worklist := make(chan []string)
    var n int // number of pending sends to worklist
    seedDomainName = crawler.getDomainName(os.Args[1])
    // Start with the command-line arguments.
    n++
    go func() { worklist <- os.Args[1:] }()


    // Crawl the web concurrently.
    seen := make(map[string]bool)

    for ; n > 0; n-- {
        for _, link := range golang_site_crawler.list {
            if !seen[link] {
                seen[link] = true
                n++
                go func(link string) {
                    worklist <- golang_site_crawler.crawl(seedDomainName, link)
                }(link)
            }
        }
    }
}  