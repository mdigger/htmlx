package htmlx

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

// AttrVal returns the attribute value with the specified key name.
// If the attribute is not specified, the false flag is returned by the second
// value.
func AttrVal(attr []html.Attribute, key string) (val string, ok bool) {
	for _, attr := range attr {
		if attr.Key == key {
			return attr.Val, true
		}
	}

	return
}

// SetAttr set the new attribute value with the specified key name.
func SetAttr(attrs []html.Attribute, key, val string) []html.Attribute {
	for _, attr := range attrs {
		if attr.Key == key {
			attr.Val = val
			return attrs
		}
	}

	attr := html.Attribute{Key: key, Val: val}
	return append(attrs, attr)
}

// RemoveAttr removes the attribute with the specified key name.
func RemoveAttr(attrs []html.Attribute, key string) []html.Attribute {
	for i, attr := range attrs {
		if attr.Key == key {
			return append(attrs[:i], attrs[i+1:]...)
		}
	}

	return attrs
}

// HasAttrWord returns true if the attribute value with the specified key name
// and specified word in value is found.
func HasAttrWord(attrs []html.Attribute, key, word string) bool {
	val, ok := AttrVal(attrs, key)
	if !ok {
		return false
	}

	for _, w := range strings.Fields(val) {
		if w == word {
			return true
		}
	}

	return false
}

// AddAttrWord add new word to attribute value.
func AddAttrWord(attrs []html.Attribute, key, word string) []html.Attribute {
	val, ok := AttrVal(attrs, key)
	if !ok || val == "" {
		return SetAttr(attrs, key, word)
	}

	for _, w := range strings.Fields(val) {
		if w == val {
			return attrs
		}
	}

	return SetAttr(attrs, key, fmt.Sprintln(val, word))
}
