package htmlx

import (
	"bytes"
	"strings"

	"github.com/mdigger/wstat"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// Rename HTML element.
func Rename(n *html.Node, name string) {
	if n == nil || n.Type != html.ElementNode || name == "" {
		return
	}

	n.Data = name
	n.DataAtom = atom.Lookup([]byte(name))
}

// Remove removes the specified element from the HTML tree.
func Remove(n *html.Node) {
	if n == nil || n.Parent == nil {
		return
	}

	n.Parent.RemoveChild(n)
}

// RemoveChilds removes all child elements if they are.
func RemoveChilds(n *html.Node) {
	if n == nil {
		return
	}

	for c := n.FirstChild; c != nil; c = n.FirstChild {
		n.RemoveChild(c)
	}
}

// HTML returns a string with HTML representation.
func HTML(n *html.Node, self bool) (string, error) {
	if n == nil {
		return "", nil
	}

	var (
		b   bytes.Buffer
		err error
	)
	if self {
		err = html.Render(&b, n)
	} else {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if err = html.Render(&b, c); err != nil {
				break
			}
		}
	}
	if err != nil {
		return "", err
	}

	return b.String(), nil
}

// SetHTML parses an HTML fragment in the context of the current element and
// replaces them the child elements.
func SetHTML(n *html.Node, data string) (err error) {
	if n == nil {
		return nil
	}

	r := strings.NewReader(data)
	nodes, err := html.ParseFragment(r, n)
	if err != nil {
		return err
	}

	RemoveChilds(n)
	for _, node := range nodes {
		n.AppendChild(node)
	}

	return nil
}

// SetText replaces the text of the element to the new one.
func SetText(n *html.Node, text string) error {
	return SetHTML(n, html.EscapeString(text))
}

// Text returns only a text representation, without HTML elements.
// Elements from the TextignoreAtom list are ignored with all the daughter
// elements.
func Text(n *html.Node) string {
	var buf strings.Builder
	_ = WriteText(&buf, n, TextIgnoreAtom)
	return buf.String()
}

// Stats returns statistics on the text.
func Stats(n *html.Node) (c wstat.Counter) {
	_ = WriteText(&c, n, TextIgnoreAtom)
	return
}

// Find finds the first coincidence on the element, including himself, and returns it.
func Find(n *html.Node, m Matcher) *html.Node {
	if n == nil || m == nil {
		return nil
	}

	var walker func(node *html.Node) *html.Node
	walker = func(node *html.Node) *html.Node {
		if node.Type == html.ElementNode && m(node) {
			return node
		}
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			if n := walker(c); n != nil {
				return n
			}
		}
		return nil
	}

	return walker(n)
}

// FindAll finds and returns all coincidences with the specified template.
func FindAll(n *html.Node, m Matcher) (result []*html.Node) {
	if n == nil || m == nil {
		return
	}

	var walker func(node *html.Node)
	walker = func(node *html.Node) {
		if node.Type == html.ElementNode && m(node) {
			result = append(result, node)
		}
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			walker(c)
		}
	}

	walker(n)
	return
}

// FindNext finds the first siblings element.
func FindNext(n *html.Node, m Matcher) *html.Node {
	if n == nil || m == nil {
		return nil
	}

	for c := n.NextSibling; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode && m(c) {
			return c
		}
	}

	return nil
}

// FindPrev finds the previous siblin element.
func FindPrev(n *html.Node, m Matcher) *html.Node {
	if n == nil || m == nil {
		return nil
	}

	for c := n.PrevSibling; c != nil; c = c.PrevSibling {
		if c.Type == html.ElementNode && m(c) {
			return c
		}
	}

	return nil
}
