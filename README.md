# htmlx

htmlx is a library which provides a set of extensions on go's `golang.org/x/net/html` library.

[![Go Reference](https://pkg.go.dev/badge/github.com/mdigger/htmlx.svg)](https://pkg.go.dev/github.com/mdigger/htmlx)

```golang
// parse html file
doc, err := htmlx.Open("test.html")
if err != nil {
    panic(err)
}

// find all link in element with id "test"
elements := doc.Find(htmlx.ID("test")).FindAll(htmlx.TagName("a"))

// print all href attribute of fined links
for _, e := range elements {
    if href, ok := e.Attr("href"); ok {
        fmt.Println(href)
    }
}
```