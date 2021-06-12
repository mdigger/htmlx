package htmlx

import "testing"

func TestNode(t *testing.T) {
	doc, err := String(`<ul id="test">
	<li><a href="test1.html">test1</a></li>
	<!-- comment -->
	<li><a href="test2.html">test2</a></li>
	<li><a href="test3.html">test3</a></li>
</ul>`)
	if err != nil {
		t.Fatal(err)
	}

	ul := doc.Find(ID("test"))
	if ul.IsEmpty() || ul.ID() != "test" {
		t.Errorf("must specify an element with the 'test' identifier")
	}

	tagLIFinder := TagName("li")
	for li := ul.Find(tagLIFinder); !li.IsEmpty(); li = li.FindNext(tagLIFinder) {
		if li.FirstChild().Data != "a" {
			t.Errorf("must specify the a element inside li")
		}
		// println(li.String())
	}

}

func TestEmpty(t *testing.T) {
	var empty Node

	if !empty.IsEmpty() {
		t.Errorf("expected node to be empty")
	}
	if empty.Parent() != empty {
		t.Errorf("expected empty.Parent to return also empty node")
	}
	if empty.FirstChild() != empty {
		t.Errorf("expected empty.FirstChild to return also empty node")
	}
	if empty.PrevSibling() != empty {
		t.Errorf("expected empty.PrevSibling to return also empty node")
	}
	if empty.NextSibling() != empty {
		t.Errorf("expected empty.NextSibling to return also empty node")
	}
	if empty.String() != "" {
		t.Errorf("expected empty.String to return empty string")
	}
	if empty.Find(ID("whatever")) != empty {
		t.Errorf("expected empty.Find to return also empty node")
	}
	if empty.FindNext(ID("whatever")) != empty {
		t.Errorf("expected empty.FindNext to return also empty node")
	}
}
