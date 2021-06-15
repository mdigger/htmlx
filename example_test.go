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

	ul := doc.Find(htmlx.ID("test"))
	for _, e := range ul.FindAll(htmlx.TagName("a")) {
		if href, ok := e.Attr("href"); ok {
			fmt.Println(href)
		}
	}

	err = ul.SetInnerHTML(`<li>no links</li>`)
	if err != nil {
		panic(err)
	}

	fmt.Println(ul)
	// Output:
	// test1.html
	// test2.html
	// test3.html
	// <ul id="test"><li>no links</li></ul>
}
