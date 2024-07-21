package main

import (
	"errors"
	"fmt"
)

var (
	ErrEmptyQueue = errors.New("empty queue")
)

func main() {
	graph := make(map[string][]string)

	graph["Ahmad"] = []string{"Steven", "Hamdy", "Yehia", "Atyea", "Hussein"}
	graph["Steven"] = []string{"Samir", "Elghamry", "Elsawy"}
	graph["Yehia"] = []string{"Samir", "Hamada", "Atef", "Alaa"}
	graph["Hamdy"] = []string{"Essam", "Ashraf", "Elghamry"}
	graph["Atyea"] = []string{"Shrshr", "Sherif", "Elsawy"}
	graph["Hussein"] = []string{"Ashmawy", "Mo2men", "Shrshr"}

	BFS(graph, "Ahmad")
}

type Node struct {
	Value string
	Next  *Node
}

type Queue struct {
	Head   *Node
	Tail   *Node
	Length int
}

func (q *Queue) Enque(value string) {
	q.Length++
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
	q.Length--
	return head.Value, nil
}

func (q *Queue) Empty() bool {
	return q.Length == 0
}

func BFS(graph map[string][]string, start string) {
	q := Queue{}
	marked := make(map[string]bool)
	q.Enque(start)

	for !q.Empty() {
		val, _ := q.Deque()
		fmt.Printf("val: %v\n", val)

		for _, child := range graph[val] {
			if !marked[child] {
				q.Enque(child)
				marked[child] = true
			} else {
				fmt.Println(child, "already marked")
			}
		} 
	}
}
