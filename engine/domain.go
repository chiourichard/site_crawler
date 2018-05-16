package engine

import (
      "log"
      "net/url"
      "strings"
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

func isSameDomain(weburl string) bool {
    if(!isValidUrl(weburl)) {
        return false
    }

    if(seedDomainName == getDomainName(weburl)) {
        return true
    } else {
        return false
    }
}