package gobt

import ()

type TreeKey struct {
	Tree string
	Key  string
}

type TreeNodeKey struct {
	Tree string
	Node string
	Key  string
}

type BBValue interface{}

type Blackboard struct {
	OpenNodes []*Node
	Globals   map[string]BBValue
	Trees     map[TreeKey]BBValue
	Nodes     map[TreeNodeKey]BBValue
}

func NewBlackboard() *Blackboard {
	b := &Blackboard{}
	b.OpenNodes = make([]*Node, 0)
	b.Globals = make(map[string]BBValue)
	b.Trees = make(map[TreeKey]BBValue)
	b.Nodes = make(map[TreeNodeKey]BBValue)
	return b
}

func (b *Blackboard) set(treeScope, nodeScope, key string, val BBValue) {
	if treeScope != "" {
		if nodeScope != "" {
			b.Nodes[TreeNodeKey{Tree: treeScope, Node: nodeScope, Key: key}] = val
			return
		}
		b.Trees[TreeKey{Tree: treeScope, Key: key}] = val
		return
	}
	b.Globals[key] = val
}

func (b *Blackboard) get(treeScope, nodeScope, key string) BBValue {
	if treeScope != "" {
		if nodeScope != "" {
			return b.Nodes[TreeNodeKey{Tree: treeScope, Node: nodeScope, Key: key}]
		}
		return b.Trees[TreeKey{Tree: treeScope, Key: key}]
	}
	return b.Globals[key]
}

func (b *Blackboard) GSet(key string, val BBValue) {
	b.set("", "", key, val)
}

func (b *Blackboard) GGet(key string) BBValue {
	return b.get("", "", key)
}

func (b *Blackboard) TSet(key, treeScope string, val BBValue) {
	b.set(treeScope, "", key, val)
}

func (b *Blackboard) TGet(key, treeScope string) BBValue {
	return b.get(treeScope, "", key)
}

func (b *Blackboard) NSet(key, treeScope, nodeScope string, val BBValue) {
	b.set(treeScope, nodeScope, key, val)
}

func (b *Blackboard) NGet(key, treeScope, nodeScope string) BBValue {
	return b.get(treeScope, nodeScope, key)
}
