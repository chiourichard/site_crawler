package engine

import (
	"testing"
)

func TestCrawlWithTrueFormatUrl(t *testing.T) {
	var webUrl = "https://www.google.com"
	SeedDomainName = GetDomainName(webUrl)
	var folderPath string = SeedDomainName 
	CreateFolder(folderPath, 0777) 
	var list []string = Crawl(webUrl)
	
	if list == nil {
		t.Errorf("func Crawl have bugs")
	}
}

func TestCrawlWithWrongFormatUrl(t *testing.T) {
	var webUrl = "test"
	SeedDomainName = GetDomainName(webUrl)
	var folderPath string = SeedDomainName 
	CreateFolder(folderPath, 0777) 
	var list []string = Crawl(webUrl)
	
	if list != nil {
		t.Errorf("func Crawl have bugs")
	}
}

func TestExtractUrlWithSameDomainOfSeed(t *testing.T) {
	var webUrl = "https://www.google.com"
	SeedDomainName = GetDomainName(webUrl)

	webUrlList, _ := Extract(webUrl)
	
	if webUrlList == nil {
		t.Errorf("func Extract have bugs")
	}
}

func TestExtractUrlWithDifferentDomainOfSeed(t *testing.T) {
	var webUrl = "https://www.google.com"
	SeedDomainName = "test.test"

	webUrlList, _ := Extract(webUrl)
	
	if webUrlList != nil {
		t.Errorf("func Extract have bugs")
	}
}