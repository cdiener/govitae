govitae
======

Render résumés with Go. Write it once and render it on the Go.

Current status
-------------

- [x] schema implemented
- [x] JSON parser working
- [x] YAML parser working
- [x] text output working
- [x] Latex output working
- [x] HTML output working

What it does
-----------

Write your cv once in JSON or YAML and have it rendered by govitae in various formats:

- a nice looking web page
- a minimal but stylish text version
- a good looking Latex version which you can compile to pdf

How to install and use
--------------------

For now you can install a development version, meaning you need to 
[install Go](https://golang.org/doc/install). After installing and setting
your $GOPATH you can get govitae with

```go
go get github.com/cdiener/govitae
go install github.com/cdiener/govitae
```

Use the `resume.json` or `minimal.json` as a starting point to build your
own and render with

```go
govitae my_cv.json
```
