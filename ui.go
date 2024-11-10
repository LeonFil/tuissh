package main

import (
	"errors"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/gdamore/tcell/v2"
	tv "github.com/rivo/tview"
)

func loadServerTree(mountPoint *tv.TreeNode, sshEntries []*SSHEntry) error {
	if mountPoint == nil {
		return errors.New("mountPoint is nil")
	}
	for _, entry := range sshEntries {
		parts := strings.Split(entry.displayName, ":")

		parent := mountPoint
		for i, part := range parts {
			var curNode *tv.TreeNode
			existNodes := parent.GetChildren()

			index := slices.IndexFunc(existNodes, func(node *tv.TreeNode) bool {
				return node.GetText() == part
			})
			if index < 0 {
				curNode = tv.NewTreeNode(part)
				parent.AddChild(curNode)
			} else {
				curNode = existNodes[index]
			}
			if i == len(parts)-1 {
				curNode.SetReference(entry)
			} else {
				curNode.SetColor(tcell.ColorBlue)
			}
			curNode.SetSelectable(true)
			parent = curNode
		}
	}
	return nil
}

func createServerTree(configFiles ...string) (*tv.TreeView, error) {
	serverTree := tv.NewTreeView()
	serverTree.SetBorder(true)
	serverTree.SetTitle("available servers")
	rootNode := tv.NewTreeNode("servers")
	serverTree.SetRoot(rootNode).SetCurrentNode(rootNode)

	for _, configFile := range configFiles {
		entries, err := loadSSHConfigFile(configFile)
		if err != nil {
			return serverTree, err
		}
		loadServerTree(rootNode, entries)
	}

	serverTree.SetChangedFunc(func(node *tv.TreeNode) {
		if len(node.GetChildren()) != 0 {
			return
		}
		// for leaf node
		serverView.SetText(fmt.Sprint(node.GetReference().(*SSHEntry)))
	})

	serverTree.SetSelectedFunc(func(node *tv.TreeNode) {
		if len(node.GetChildren()) != 0 {
			return
		}
		// for leaf node
		entry := node.GetReference().(*SSHEntry)
		msg := errView.GetText(true)
		msg += logMsg(fmt.Sprintf("connecting to %s...", entry.displayName))
		errView.SetText(msg)
		errView.ScrollToEnd()

		// run command
		app.Suspend(func() {
			// at this time, the terminal is used by this function
			sshShell(entry)
		})
	})

	return serverTree, nil
}

func logMsg(msg string) string {
	now := time.Now()
	return fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d %s\n",
		now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second(), msg)
}
