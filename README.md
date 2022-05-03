# boot4go-pathmatcher
a path matcher toolkit for golang as antPathMatcher in java

![license](https://img.shields.io/badge/license-Apache--2.0-green.svg)

# Introduce

In golang, there are two built-in packages: path and filepath. The match function is provided for path matching, but the function is too weak. You can refer to the functions of path and filepath by yourself. When analyzing the path of gateway, it feels that the function can not meet the requirements at all, so I have implemented it with reference to the function of antpathmatcher.


# Feature

- Support **
- Support ^
- Support multiple square brackets

# Usage

- Add boot4go-pathmatcher with the following import
```
import pathmatcher "github.com/gohutool/boot4go-pathmatcher"
```

- Support **

```
=== RUN   TestPathMatch
pathmatcher.Match("**/b", "/b")=true
pathmatcher.Match("**/b", "/b/aaa/a.log")=false
pathmatcher.Match("**/b/**", "/b/aaa/bbb/a.log")=true
pathmatcher.Match("**/b/**/", "/b/aaa/bbb/ccc/a.log")=false
pathmatcher.Match("**/b/**/", "/b/aaa/bbb/ccc/dddd")=false
pathmatcher.Match("**/b/**", "/b/aaa/bbb/ccc/dddd")=true
pathmatcher.Match("**/b/**", "/b/aaa/bbb/ccc/dddd")=true
pathmatcher.Match("**/b/**", "/b/aaa/bbb/ccc/dddd/a.log")=true
pathmatcher.Match("**/b/**/*.log", "/b/aaa/bbb/a.log")=true`
```

- Support ^

```
== RUN   TestPathMatch
pathmatcher.Match("[^a]bc", "abc")=false
pathmatcher.Match("[^a]bc", "bbc")=true
pathmatcher.Match("[^a]b[^c]", "xbc")=false
pathmatcher.Match("[^a]b[^c]", "xbx")=true
```

- Support multiple square brackets

```
== RUN   TestPathMatch
pathmatcher.Match("[^a-c]bc", "abc")=false
pathmatcher.Match("[^a-c]bc", "bbc")=false
pathmatcher.Match("[^a-c]bc", "cbc")=false
pathmatcher.Match("[^a-c]bc", "dbc")=true
pathmatcher.Match("[^a-c]bc", "ebc")=true
pathmatcher.Match("[^a-c]b[c-d]", "dbc")=true
pathmatcher.Match("[^a-c]b[c-d]", "dbd")=true
pathmatcher.Match("[^a-c]b[c-d]", "dbe")=false
pathmatcher.Match("[^a-c]b[c-d]", "dbf")=false
```

- Support advance feature

```
== RUN   TestPathMatch
pathmatcher.Match("[a-cx-z]", "a")=true
pathmatcher.Match("[a-cx-z]", "c")=true
pathmatcher.Match("[a-cx-z]", "x")=true
pathmatcher.Match("[a-cx-z]", "z")=true
pathmatcher.Match("[a-cx-z]", "o")=false
pathmatcher.Match("[a-cx-z]", "f")=false
```

```
pathmatcher.Match("[^a-cx-z]", "a")=false
pathmatcher.Match("[^a-cx-z]", "x")=false
pathmatcher.Match("[^a-cx-z]", "o")=true
```

# Related project

- log4go https://github.com/gohutool/log4go
- boot4go-util https://github.com/gohutool/boot4go-util