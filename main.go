package main  

import (
	"os"
	"github.com/chiourichard/site_crawler/engine"
) 

func main() {
	worklist := make(chan []string)
	var numWorklist int 
	engine.SeedDomainName = engine.GetDomainName(os.Args[1])
	// Start with the command-line arguments.
	numWorklist++
	go func() { worklist <- os.Args[1:] }()

	var folderPath string = engine.SeedDomainName 
	err := engine.CreateFolder(folderPath, 0777) 
	if err != nil  {
		return 
	} else {
		engine.FolderName = folderPath
	}

	seen := make(map[string]bool)

	for ; numWorklist > 0; numWorklist-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				numWorklist++
				go func(link string) {
					worklist <- engine.Crawl(link)
				}(link)
			}
		}
	}
}  
