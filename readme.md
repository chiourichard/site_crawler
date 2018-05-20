A site crawler tool that will crawl all pages with same domain name of a site to your computer:

<!-- MarkdownTOC -->

* [Todo List](#todo-list)
* [Installing](#installing) - [Requirements](#requirements)
* [Quickstart](#quickstart)
* [Run it](#run-it)

<!-- /MarkdownTOC -->

## Todo List

* Web-based interface.

## Installing

### Requirements

* [Go](https://golang.org/dl/) 1.6+

## Quickstart

Install goreporter (see above).

```
$ go get -u github.com/chiourichard/site_crawler
$ go build
```

## Run it:

It's a command line tool:

```
$ ./site_crawler https://www.google.com
```

Then you can see a new folder created named `google.com` in the path you are in now. You can find the pages this tool crawl for you.
