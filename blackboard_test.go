package gobt

import (
	"testing"
)

func TestBlackboard(t *testing.T) {
	var r BBValue
	bb := NewBlackboard()
	if len(bb.OpenNodes) != 0 {
		t.Error("OpenNodes should be empty")
	}

	bb.GSet("Foo", 1)
	r = bb.GGet("Foo")
	if r != 1 {
		t.Error("global was not set properly. expected 1, got ", r)
	}

	bb.TSet("tree1", "Foo", 2)
	r = bb.TGet("tree1", "Foo")
	if r != 2 {
		t.Error("tree scope was not set properly. expected 2, got ", r)
	}

	bb.NSet("tree1", "node1", "Foo", 3)
	r = bb.NGet("tree1", "node1", "Foo")
	if r != 3 {
		t.Error("node scope was not set properly. expected 3, got ", r)
	}
}
