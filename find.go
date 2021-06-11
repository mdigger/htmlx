package htmlx

import "golang.org/x/net/html"

// Matcher used as synonym the functions for searching and selecting HTML elements.
type Matcher = func(*html.Node) bool

// Find finds the first coincidence on the element, including himself, and returns it.
func Find(n *html.Node, m Matcher) *html.Node {
	if n == nil || m == nil {
		return nil
	}

	var walker func(node *html.Node) *html.Node
	walker = func(node *html.Node) *html.Node {
		if m(node) {
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
		if m(node) {
			result = append(result, node)
		}

		for c := node.FirstChild; c != nil; c = c.NextSibling {
			walker(c)
		}

	}

	walker(n)
	return
}

// FindSibling finds the first coincidence on the element that follows the
// specified one. When searching does not use recursion and is not compared with
// the current element.
func FindSibling(n *html.Node, m Matcher) *html.Node {
	if n == nil || m == nil {
		return nil
	}

	for c := n.NextSibling; c != nil; c = c.NextSibling {
		if m(c) {
			return c
		}
	}

	return nil
}

// FindPrevSibling finds the previous siblin element that follows the
// specified one. When searching does not use recursion and is not compared with
// the current element.
func FindPrevSibling(n *html.Node, m Matcher) *html.Node {
	if n == nil || m == nil {
		return nil
	}

	for c := n.PrevSibling; c != nil; c = c.PrevSibling {
		if m(c) {
			return c
		}
	}

	return nil
}
