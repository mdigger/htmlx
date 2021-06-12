package htmlx_test

import (
	"fmt"
	"strings"

	"github.com/mdigger/htmlx"
)

func Example() {
	source := `<ul id="test">
	<li><a href="test1.html">test1</a></li>
	<li><a href="test2.html">test2</a></li>
	<li><a href="test3.html">test3</a></li>
</ul>`
	doc, err := htmlx.Parse(strings.NewReader(source))
	if err != nil {
		panic(err)
	}

	elements := doc.Find(htmlx.ID("test")).FindAll(htmlx.TagName("a"))
	for _, e := range elements {
		if href, ok := e.Attr("href"); ok {
			fmt.Println(href)
		}
	}
	// Output:
	// test1.html
	// test2.html
	// test3.html
}
