package htmlx

import "testing"

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
