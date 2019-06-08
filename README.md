# goexpand

[![Build Status](https://travis-ci.org/takebayashi/goexpand.svg?branch=master)](https://travis-ci.org/takebayashi/goexpand)

```:go
goexpand.Expand("foo[0:1][2:3],bar[00:01]")
// => {"foo02", "foo03", "foo12", "foo13", "bar00", "bar01"}
```
