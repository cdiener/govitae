govitae
======

Render JSON résumés with Go.

Current status
-------------

- [x] schema implemented
- [x] parser working
- [ ] text output working
- [ ] HTML output working
- [ ] Latex output working

What it does
-----------

Write your cv once in JSON and have it rendered by govitae in various formats:

- a nice looking web page
- a minimal but stylish text version
- a good looking Latex version which you can compile to pdf

How to install and use
--------------------

```go
go get github.com/cdiener/govitae
```

Use the `resume.json` or `minimal.json` as a starting point and render with

```go
govitae my_cv.json
```
