package engine

import (
	"testing"
)

func TestGetDomainNameWithRegularUrl(t *testing.T) {
	var webUrl = "https://www.google.com"
	var testDomainName string = GetDomainName(webUrl)

	if testDomainName != "google.com" {
		t.Errorf("func GetDomainName can't get url domain")
	}
}

func TestGetDomainNameWithFalseUrl(t *testing.T) {
	var webUrl = "javascript();"
	var testDomainName string = GetDomainName(webUrl)

	if testDomainName != "" {
		t.Errorf("func GetDomainName can't get url domain")
	}
}

func TestSameDomainNameWithSeedDomain(t *testing.T) {
	var webUrl = "https://www.google.com"
	SeedDomainName = "google.com"

	if !IsSameDomain(webUrl) {
		t.Errorf("func IsSameDomain have bugs")
	}
}

func TestDifferentDomainNameWithSeedDomain(t *testing.T) {
	var webUrl = "https://www.google.com"
	SeedDomainName = "kkbox.com"

	if IsSameDomain(webUrl) {
		t.Errorf("func IsSameDomain can't distinguish different domains")
	}
}
func TestWrongUrlWillFail(t *testing.T) {
	var webUrl = "ftp://www.google.com"

	if IsValidUrl(webUrl) {
		t.Errorf("func IsValidUrl can't distinguish a url is invalid.")
	}
}

func TestRightUrlWillSuccess(t *testing.T) {
	var webUrl = "https://tw.yahoo.com/"

	if !IsValidUrl(webUrl) {
		t.Errorf("func IsValidUrl can't distinguish a url is valid")
	}
}
