package htmlx

import (
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// Matcher used as synonym the functions for searching and selecting HTML elements.
type Matcher = func(*html.Node) bool

// Tag used to search for items on a tag name identifier.
func Tag(a atom.Atom) Matcher {
	return func(n *html.Node) bool {
		return n.DataAtom == a
	}
}

// Tag used to search for items on a tag name.
func TagName(name string) Matcher {
	if a := atom.Lookup([]byte(name)); a != 0 {
		return Tag(a)
	}

	return func(n *html.Node) bool {
		return n.Data == name
	}
}

// HasAttrVal is used to find an element with a specified attribute value.
func HasAttrVal(name, value string) Matcher {
	return func(n *html.Node) bool {
		val, ok := AttrVal(n.Attr, name)
		return ok && val == value
	}
}

// ID is used to find an element with a specified unique identifier.
func ID(id string) Matcher {
	return HasAttrVal(AttrID, id)
}

// Class is used to select elements with a specified style class.
func Class(name string) Matcher {
	return func(n *html.Node) bool {
		return HasAttrWord(n.Attr, AttrClass, name)
	}
}
