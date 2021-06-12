# htmlx

htmlx is a library which provides a set of extensions on go's `golang.org/x/net/html` library.

[![Go Reference](https://pkg.go.dev/badge/github.com/mdigger/htmlx.svg)](https://pkg.go.dev/github.com/mdigger/htmlx)

```golang
// parse html file
doc, err := htmlx.Open("test.html")
if err != nil {
    panic(err)
}

ul := doc.Find(htmlx.ID("test"))
for _, e := range ul.FindAll(htmlx.TagName("a")) {
    if href, ok := e.Attr("href"); ok {
        fmt.Println(href)
    }
}

err = ul.SetHTML(`<li>no links</li>`)
if err != nil {
    panic(err)
}
fmt.Println("html:", ul)
```