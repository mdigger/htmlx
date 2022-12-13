package htmlx

import (
	"io"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// WriteText walk all the invested text nodes and records the text from them
// to the specified StringWriter. Ignore comments.
//
// BUG: <noscript> parsed as text.
func WriteText(w io.StringWriter, n *html.Node, ignore map[atom.Atom]bool) error {
	if n == nil {
		return nil
	}

	var output func(*html.Node) error
	output = func(n *html.Node) (err error) {
		switch n.Type {
		case html.TextNode:
			_, err = w.WriteString(
				html.UnescapeString(n.Data))
			return err
		case html.CommentNode:
			return nil
		case html.ElementNode:
			if ignore[n.DataAtom] {
				return nil
			}
		}

		for child := n.FirstChild; child != nil; child = child.NextSibling {
			if err = output(child); err != nil {
				return err
			}
		}

		return nil
	}

	return output(n)
}

// TextIgnoreAtom specifies the list of items whose contents are ignored when
// working with text nodes.
var TextIgnoreAtom = map[atom.Atom]bool{
	// ignore text inside
	atom.Head:   true,
	atom.Script: true,
	atom.Style:  true,
	// self closed (optimization)
	atom.Area:   true,
	atom.Base:   true,
	atom.Br:     true,
	atom.Col:    true,
	atom.Embed:  true,
	atom.Hr:     true,
	atom.Img:    true,
	atom.Input:  true,
	atom.Keygen: true,
	atom.Link:   true,
	atom.Meta:   true,
	atom.Param:  true,
	atom.Source: true,
	atom.Track:  true,
	atom.Wbr:    true,
}
