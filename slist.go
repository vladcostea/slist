package slist

import (
	"errors"
)

type NodeItem interface {
	Cmp(key NodeItem) bool
	Eq(key NodeItem) bool
}

type Node struct {
	item     NodeItem
	nextNode *Node
}

func (n *Node) Next() *Node {
	return n.nextNode
}

type onEqualHandler func(node *Node, item NodeItem)

type SList struct {
	head    *Node
	onEqual onEqualHandler
}

func (l *SList) Head() *Node {
	return l.head
}

func (l *SList) OnEqual(handler onEqualHandler) {
	l.onEqual = handler
}

func onEqual(node *Node, item NodeItem) {
}

func NewSList() *SList {
	return &SList{onEqual: onEqual}
}

func (l *SList) IsEmpty() bool {
	return l.head == nil
}

func (l *SList) Insert(item NodeItem) {
	if l.head == nil {
		l.head = &Node{item: item, nextNode: nil}
		return
	}

	var found bool

	current := l.head
	previous := l.head
	newNode := &Node{item: item, nextNode: nil}

	for {
		if current.item.Cmp(item) {
			if current == l.head {
				newNode.nextNode = current
				l.head = newNode
			} else {
				newNode.nextNode = previous.nextNode
				previous.nextNode = newNode
			}
			found = true
			break
		} else if current.item.Eq(item) {
			l.onEqual(current, item)
			found = true
			break
		}

		if current.nextNode == nil {
			break
		}

		previous = current
		current = current.nextNode
	}

	if found == false {
		current.nextNode = newNode
	}
}

func (l *SList) Find(item NodeItem) NodeItem {
	if l.IsEmpty() {
		return nil
	}

	for current := l.head; current.nextNode != nil; current = current.nextNode {
		if current.item.Eq(item) {
			return current.item
		}
	}

	return nil
}

func (l *SList) Remove(item NodeItem) (NodeItem, error) {
	if l.IsEmpty() {
		return nil, errors.New("Item not found")
	}

	previous := l.head
	current := l.head

	for current.nextNode != nil {
		if current.item.Eq(item) {
			if current == l.head {
				l.head = l.head.nextNode
			} else {
				previous.nextNode = current.nextNode
			}

			return current.item, nil
		}

		previous = current
		current = current.nextNode
	}

	return nil, errors.New("Item not found")
}

func (l *SList) Log(logger func(item NodeItem)) {
	current := l.head
	for current != nil {
		logger(current.item)
		current = current.nextNode
	}
}
