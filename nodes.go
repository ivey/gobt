package gobt

import ()

type NodeStatus int

const (
	_ NodeStatus = iota
	SUCCESS
	FAILURE
	RUNNING
	ERROR
)

type Node interface {
	name() string
	tick(*Walker) NodeStatus
	setChildren([]Node)
}

// Sequence
type Sequence struct {
	Name     string
	children []Node
}

func (n *Sequence) name() string {
	return n.Name
}

func (n *Sequence) setChildren(c []Node) {
	n.children = c
}

func (n *Sequence) tick(w *Walker) NodeStatus {
	for _, child := range n.children {
		// TODO: if we had a memorized place, skip to that
		status := w.walkNode(child)
		if status != SUCCESS {
			// TODO: if RUNNING, memorize our place
			return status
		}
	}
	return SUCCESS
}

// Selector
type Selector struct {
	Name     string
	children []Node
}

func (n *Selector) name() string {
	return n.Name
}

func (n *Selector) setChildren(c []Node) {
	n.children = c
}

func (n *Selector) tick(w *Walker) NodeStatus {
	for _, child := range n.children {
		// TODO: if we had a memorized place, skip to that
		status := w.walkNode(child)
		if status != FAILURE {
			// TODO: if RUNNING, memorize our place
			return status
		}
	}
	return FAILURE
}

// Inverter
type Inverter struct {
	Name     string
	children []Node
}

func (n *Inverter) name() string {
	return n.Name
}

func (n *Inverter) setChildren(c []Node) {
	n.children = c
}

func (n *Inverter) tick(w *Walker) NodeStatus {
	if len(n.children) != 1 {
		return FAILURE
	}
	status := w.walkNode(n.children[0])
	if status == FAILURE {
		return SUCCESS
	}
	if status == SUCCESS {
		return FAILURE
	}
	return status
}

// Delayer

// Action
type Action struct {
	Name   string
	action func(interface{}, *Blackboard) bool
}

func (n *Action) name() string {
	return n.Name
}

func (n *Action) setChildren(c []Node) {
	// noop
}

func (n *Action) tick(w *Walker) NodeStatus {
	if n.action != nil {
		if n.action(w.target, w.blackboard) != true {
			return FAILURE
		}
	}
	return SUCCESS
}
