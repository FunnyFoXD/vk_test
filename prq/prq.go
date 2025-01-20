package prq

import (
	"githab.com/funnyfoxd/vk_test/models"
)

type PriorityQueue []*models.Node

func (p PriorityQueue) Len() int {
	return len(p)
}

func (p PriorityQueue) Less(i, j int) bool {
	return p[i].FCost < p[j].FCost
}

func (p PriorityQueue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *PriorityQueue) Push(x interface{}) {
	*p = append(*p, x.(*models.Node))
}

func (p *PriorityQueue) Pop() interface{} {
	old := *p
	n := len(old)
	item := old[n-1]
	*p = old[0 : n-1]

	return item
}