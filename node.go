package htmlx

import (
	"io"
	"net/http"
	"os"

	"github.com/mdigger/wstat"
	"golang.org/x/net/html"
	"golang.org/x/net/html/charset"
)

// Node expands html.node with additional methods.
type Node struct {
	*html.Node
}

// New warps the representation of html.node by adding it a new functionality.
func New(node *html.Node) Node {
	return Node{Node: node}
}

// Parse returns a parsed HTML tree representation.
func Parse(r io.Reader) (n Node, err error) {
	doc, err := html.Parse(r)
	return New(doc), err
}

// Load loads and parses an HTML document from the file.
func Load(path string) (Node, error) {
	f, err := os.Open(path)
	if err != nil {
		return New(nil), err
	}
	defer f.Close()

	return Parse(f)
}

// Get loads and parses an HTML document at the specified url address.
func Get(url string) (Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return New(nil), err
	}
	defer resp.Body.Close()

	r, err := charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))
	if err != nil {
		return New(nil), err
	}

	return Parse(r)
}

// IsEmpty returns true if the node is not specified.
func (n Node) IsEmpty() bool {
	return n.Node == nil
}

// Rename renames HTML element.
func (n *Node) Rename(name string) {
	Rename(n.Node, name)
}

// Remove removes the specified element from the HTML tree.
func (n *Node) Remove() {
	Remove(n.Node)
}

// RemoveChilds removes all child elements if they are exists.
func (n *Node) RemoveChilds() {
	RemoveChilds(n.Node)
}

// Parent returns the parent element.
func (n Node) Parent() Node {
	if n.IsEmpty() {
		return n
	}
	return New(n.Node.Parent)
}

// FirstChild returns the first child element.
func (n Node) FirstChild() Node {
	if n.IsEmpty() {
		return n
	}
	return New(n.Node.FirstChild)
}

// LastChild returns the last child element.
func (n Node) LastChild() Node {
	if n.IsEmpty() {
		return n
	}
	return New(n.Node.LastChild)
}

// PrevSibling returns the previous sibling element.
func (n Node) PrevSibling() Node {
	if n.IsEmpty() {
		return n
	}
	return New(n.Node.PrevSibling)
}

// NextSibling returns the previous sibling element.
func (n Node) NextSibling() Node {
	if n.IsEmpty() {
		return n
	}
	return New(n.Node.NextSibling)
}

// Find returns the first element that is suitable for the specified conditions.
func (n Node) Find(m Matcher) Node {
	return New(Find(n.Node, m))
}

// FindAll returns all the elements suitable for the specified conditions.
func (n Node) FindAll(m Matcher) []Node {
	nodes := FindAll(n.Node, m)
	if len(nodes) == 0 {
		return nil
	}

	result := make([]Node, len(nodes))
	for i, node := range nodes {
		result[i] = New(node)
	}

	return result
}

// FindNext finds the first siblin element.
func (n Node) FindNext(m Matcher) Node {
	return New(FindNext(n.Node, m))
}

// FindPrev finds the previous siblin element.
func (n Node) FindPrev(m Matcher) Node {
	return New(FindPrev(n.Node, m))
}

// HTML returns a string with HTML representation.
func (n Node) HTML() (string, error) {
	return HTML(n.Node)
}

// String returns a string with HTML representation.
// Possible error is ignored.
func (n Node) String() string {
	str, _ := n.HTML()
	return str
}

// SetHTML parses an HTML fragment in the context of the current element and
// replaces them the child elements.
func (n *Node) SetHTML(data string) error {
	return SetHTML(n.Node, data)
}

// Text returns only a text representation, without HTML elements.
func (n Node) Text() string {
	return Text(n.Node)
}

// SetText replaces the text of the element to the new and removes possible
// child items.
func (n *Node) SetText(text string) error {
	return SetText(n.Node, text)
}

// Stats returns statistics on the text.
func (n Node) Stats() (c wstat.Counter) {
	return Stats(n.Node)
}

// Predefined attribute names.
const (
	AttrID    = "id"
	AttrClass = "class"
)

// Attr returns the attribute value with the specified name and the flag that
// the attribute was specified for this item.
func (n Node) Attr(name string) (val string, ok bool) {
	if n.IsEmpty() || n.Type != html.ElementNode {
		return
	}

	return AttrVal(n.Node.Attr, name)
}

// ID returns the unique identifier of the element.
func (n Node) ID() string {
	val, _ := n.Attr(AttrID)
	return val
}

// SetAttr set the new attribute value with the specified name.
func (n *Node) SetAttr(name, value string) {
	if n.IsEmpty() || n.Type != html.ElementNode {
		return
	}

	n.Node.Attr = SetAttr(n.Node.Attr, name, value)
}

// RemoveAttr removes the attribute value with the specified name.
func (n *Node) RemoveAttr(name, value string) {
	if n.IsEmpty() || n.Type != html.ElementNode {
		return
	}

	n.Node.Attr = RemoveAttr(n.Node.Attr, name)
}

// HasClass returns true if the item is specified with the specified name.
func (n Node) HasClass(name string) (ok bool) {
	if n.IsEmpty() || n.Type != html.ElementNode {
		return
	}

	return HasAttrWord(n.Node.Attr, AttrClass, name)
}

// AddClass adds a new style name to the element attribute list.
func (n *Node) AddClass(name string) {
	if n.IsEmpty() || n.Type != html.ElementNode {
		return
	}

	n.Node.Attr = AddAttrWord(n.Node.Attr, AttrClass, name)
}
