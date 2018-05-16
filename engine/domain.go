package engine

import (
      "log"
      "net/url"
      "strings"
      "regexp"
)

func getDomainName(weburl string) string {
      u, err := url.Parse(weburl)
      if err != nil {
              log.Fatal(err)
      }
      parts := strings.Split(u.Hostname(), ".")

      if(len(parts) > 1) {
        domain := parts[len(parts)-2] + "." + parts[len(parts)-1]
        return domain
      } else {
        return ""
      }
}

func isSameDomain(seedDomainName string, weburl string) bool {
    if(!isValidUrl(weburl)) {
        return false
    }

    if(seedDomainName == getDomainName(weburl)) {
        return true
    } else {
        return false
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