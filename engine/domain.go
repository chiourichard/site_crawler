package engine

import (
	"log"
	"net/url"
	"strings"
	"regexp"
)

var SeedDomainName string = ""

func GetDomainName(weburl string) string {
	u, err := url.Parse(weburl)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	parts := strings.Split(u.Hostname(), ".")

	if(len(parts) > 1) {
		domain := parts[len(parts)-2] + "." + parts[len(parts)-1]
		return domain
	} else {
		return ""
	}
}

func IsSameDomain(weburl string) bool {
	if(!IsValidUrl(weburl)) {
		return false
	}

	if(SeedDomainName == GetDomainName(weburl)) {
		return true
	} else {
		return false
	}
}

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
