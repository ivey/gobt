package gobt

import ()

type Tree struct {
	Root Node
	Name string
}

type Walker struct {
	target     interface{}
	blackboard *Blackboard
	openNodes  []Node
	nodeCount  int
	treeName   string
}

func (w *Walker) isOpen(nodeName string) bool {
	return w.blackboard.NGet(w.treeName, nodeName, "open") != true
}

func (w *Walker) open(node Node) {
	w.openNodes = append(w.openNodes, node)
	w.nodeCount++
	w.blackboard.NSet(w.treeName, node.name(), "open", true)
}

func (w *Walker) close(node Node) {
	w.blackboard.NSet(w.treeName, node.name(), "open", false)
	w.openNodes = w.openNodes[:len(w.openNodes)-1]
}

func (w *Walker) walkNode(node Node) NodeStatus {
	// Open
	if !w.isOpen(node.name()) {
		w.open(node)
	}

	// Tick
	status := node.tick(w)

	// Close
	if status != RUNNING {
		w.close(node)
	}

	return status
}

func (w *Walker) stillOpen(node Node) bool {
	for _, n := range w.openNodes {
		if n == node {
			return true
		}
	}
	return false
}

func newWalker(treeName string, target interface{}, blackboard *Blackboard) *Walker {
	w := &Walker{treeName: treeName, target: target, blackboard: blackboard, nodeCount: 0}
	w.openNodes = make([]Node, 0)
	return w
}

func (t *Tree) Tick(target interface{}, blackboard *Blackboard) {
	walker := newWalker(t.Name, target, blackboard)
	walker.walkNode(t.Root)

	openNodes := blackboard.TGet(t.Name, "openNodes")
	if openNodes != nil {
		for _, node := range openNodes.([]Node) {
			if !walker.stillOpen(node) {
				walker.close(node)
			}
		}
	}
}
