package main  
  
import (
    "os"
    "github.com/chiourichard/golang_site_crawler/engine"
) 

func main() {
    worklist := make(chan []string)
    var n int // number of pending sends to worklist
    engine.SeedDomainName = engine.GetDomainName(os.Args[1])
    // Start with the command-line arguments.
    n++
    go func() { worklist <- os.Args[1:] }()

    var folderPath string = engine.SeedDomainName 
    _, err := engine.CreateFolder(folderPath, 0777) 
    if err != nil  {
        return err
    } else {
        engine.FolderName = folderPath
    }

    // Crawl the web concurrently.
    seen := make(map[string]bool)

    for ; n > 0; n-- {
        webUrlList := <-worklist
        for _, webUrl := range webUrlList {
            if !seen[webUrl] {
                seen[webUrl] = true
                n++
                go func(link string) {
                    worklist <- engine.Crawl(webUrl)
                }(webUrl)
            }
        }
    }
}  