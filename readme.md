![site_crawler]

[![Build Status](https://travis-ci.org/chiourichard/site_crawler.svg?branch=master)](https://travis-ci.org/chiourichard/site_crawler)
[![GoDoc](https://godoc.org/github.com/chiourichard/site_crawler?status.svg)](https://godoc.org/github.com/chiourichard/site_crawler)

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

* [Go](https://golang.org/dl/) 1.8+

## Quickstart

Install site_crawler (see above).

```
$ go get -u github.com/chiourichard/site_crawler
$ go build
```

## Run it:

It's a command line tool:

```
$ ./site_crawler https://www.google.com
```

Then you can see a new folder named `google.com` in the path you are in now. You can find the pages this tool collects for you.

### Web service version:

In the folder "webservice", you can build web-based site_crawler:

```
$ go build
```

and execute:

```
./webservice
```

This service listen 9090, you can use browser to access it.
parameter: `url`
For example:

```
http://localhost:9090/?url=http://www.gopl.io/
```
