package main

import (
	"testing"

	"github.com/rivo/tview"
)

func Test_loadServerTree(t *testing.T) {
	if r := loadServerTree(nil, nil); r == nil {
		t.Fail()
	}
	tnode := tview.NewTreeNode("servers")
	if r := loadServerTree(tnode, nil); r != nil {
		t.Fail()
	}

	entries := []*SSHEntry{
		&SSHEntry{displayName: "a:b:c:server"},
		&SSHEntry{displayName: "a:b:server"},
		&SSHEntry{displayName: "a:server"},
		&SSHEntry{displayName: "server"},
	}

	loadServerTree(tnode, entries)
	N := 2
	if n := len(tnode.GetChildren()); n != N {
		t.Fatalf("%d != %d\n", n, N)
	}
}
