package engine

import (
	"net/url"
	"os"
	"testing"
)

func TestCreateFolderSuccess(t *testing.T) {
	var filepath string = "test"
	var mode os.FileMode = 0777

	if err := CreateFolder(filepath, mode); err != nil {
		t.Errorf("func Createfolder failed: %s", err)
	}
	os.RemoveAll(filepath)
}

func TestCreateFolderFail(t *testing.T) {
	var filepath string = "abc"
	var mode os.FileMode = 0777

	CreateFolder(filepath, mode)
	err := CreateFolder(filepath, mode)
	if err == nil {
		t.Errorf("func Createfolder failed: %s", err)
	}
}

func TestDownloadFileSuccess(t *testing.T) {
	var webUrl string = "http://www.google.com"
	var testDomainName string = GetDomainName(webUrl)

	var folderName = testDomainName
	var mode os.FileMode = 0777

	if err := CreateFolder(folderName, mode); err != nil {
		t.Errorf("func Createfolder failed: %v", err)
	}
	var filepath string = folderName + "/" + url.PathEscape(webUrl)

	err := DownloadFile(filepath, webUrl)
	if err != nil {
		t.Errorf("func Download failed: %s", err)
	}

	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		t.Errorf("download file is not exist")
	}
	os.RemoveAll(folderName)
}

func TestDownloadFileFailWithWrongPath(t *testing.T) {
	var webUrl string = "http://www.google.com"
	var testDomainName string = GetDomainName(webUrl)

	var folderName = testDomainName
	var mode os.FileMode = 0777

	if err := CreateFolder(folderName, mode); err != nil {
		t.Errorf("func Createfolder failed: %v", err)
	}
	var filepath string = "Wrongfolder" + "/" + url.PathEscape(webUrl)

	err := DownloadFile(filepath, webUrl)
	if err == nil {
		t.Errorf("func Download failed: %s", err)
	}
}

func TestDownloadFileFailWithWrongUrl(t *testing.T) {
	var webUrl string = "ftp://www.google.com"
	var testDomainName string = GetDomainName(webUrl)

	var folderName = testDomainName
	var mode os.FileMode = 0777

	if err := CreateFolder(folderName, mode); err != nil {
		t.Errorf("func Createfolder failed: %v", err)
	}
	var filepath string = folderName + "/" + url.PathEscape(webUrl)

	err := DownloadFile(filepath, webUrl)
	if err == nil {
		t.Errorf("func Download failed: %s", err)
	}
}
func TestDownloadFileFailWithFailStatus(t *testing.T) {
	var webUrl string = "https://www.google.com.tw/abc"
	var testDomainName string = GetDomainName(webUrl)

	var folderName = testDomainName
	var mode os.FileMode = 0777

	if err := CreateFolder(folderName, mode); err != nil {
		t.Errorf("func Createfolder failed: %v", err)
	}
	var filepath string = folderName + "/" + url.PathEscape(webUrl)

	err := DownloadFile(filepath, webUrl)
	if err == nil {
		t.Errorf("func Download failed: %s", err)
	}
}
