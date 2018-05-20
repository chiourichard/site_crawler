package engine

import (
	"log"
	"net/url"
	"regexp"
	"strings"
)

var SeedDomainName string = ""

// Get domain name from a web url
func GetDomainName(weburl string) string {
	u, err := url.Parse(weburl)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	parts := strings.Split(u.Hostname(), ".")

	if len(parts) > 1 {
		domain := parts[len(parts)-2] + "." + parts[len(parts)-1]
		return domain
	} else {
		return ""
	}
}

// Compare domain of a web url with seed domain
func IsSameDomain(weburl string) bool {
	if !IsValidUrl(weburl) {
		return false
	}

	if SeedDomainName == GetDomainName(weburl) {
		return true
	} else {
		return false
	}
}

// distinguish a url is vaild or not
func IsValidUrl(toTest string) bool {
	var validHttp = regexp.MustCompile("^(http|https)://")
	var matchHttp = validHttp.MatchString(toTest)

	_, err := url.ParseRequestURI(toTest)
	if err != nil || matchHttp == false {
		return false
	} else {
		return true
	}
}
