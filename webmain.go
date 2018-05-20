package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/chiourichard/site_crawler/engine"
)

func crawler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	var seedUrl string = ""
	for k, v := range r.Form {
		if k == "url" {
			seedUrl = strings.Join(v, "")
		}
	}

	if seedUrl != "" {
		worklist := make(chan []string)
		var numWorklist int
		engine.SeedDomainName = engine.GetDomainName(seedUrl)
		// Start with the command-line arguments.
		numWorklist++
		go func() { worklist <- r.Form["url"] }()

		var folderPath string = engine.SeedDomainName
		err := engine.CreateFolder(folderPath, 0777)
		if err != nil {
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

		result, _ := json.Marshal(CrawledUrls)
		fmt.Fprintf(w, string(result))
	}
}

func main() {
	http.HandleFunc("/", crawler)            //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
