package main

import "errors"

var (
	ErrEmptyQueue = errors.New("empty queue")
)

type Node struct {
	Value string
	Next *Node
}

type Queue struct {
	Head *Node
	Tail *Node
}

func (q *Queue) Enque(value string) {
	if q.Head == nil {
		q.Head = &Node{Value: value}
		q.Tail = q.Head
	} else {
		q.Tail.Next = &Node{Value: value}
		q.Tail = q.Tail.Next
	}
}

func (q *Queue) Front() (string, error) {
	if q.Head == nil {
		return "", ErrEmptyQueue
	}
	
	return q.Head.Value, nil
}

func (q *Queue) Deque() (string, error) {
	if q.Head == nil {
		return "", ErrEmptyQueue
	}

	head := q.Head
	q.Head = q.Head.Next

	return head.Value, nil
}