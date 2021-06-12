package htmlx

import "testing"

func TestEmpty(t *testing.T) {
	var empty Node

	if !empty.IsEmpty() {
		t.Errorf("expected finder to be empty")
	}
	if empty.Find(ID("whatever")) != empty {
		t.Errorf("expected empty.Find to return also empty finder")
	}
	if empty.FindSibling(ID("whatever")) != empty {
		t.Errorf("expected empty.FindSibling to return also empty finder")
	}
	if empty.FirstChild() != empty {
		t.Errorf("expected empty.FirstChild to return also empty finder")
	}
	if empty.NextSibling() != empty {
		t.Errorf("expected empty.NextSibling to return also empty finder")
	}
	if empty.String() != "" {
		t.Errorf("expected empty.String to return empty string")
	}
}
