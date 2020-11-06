package semanticVersioning

import inter "npmExtension/dependenciesTreeGenerator"

type queue struct {
	slice	[]*inter.PackageNode
}

func NewQueue() queue {
	q := queue{}
	q.slice = make([]*inter.PackageNode, 0)
	return q
}

func (q *queue) Enqueue(package_ *inter.PackageNode){
	q.slice = append(q.slice, package_)
}

func (q *queue) Dequeue() *inter.PackageNode {
	element := q.slice[0]
	q.slice = q.slice[1:]
	return element
}

func (q queue)IsEmpty() bool {
	return len(q.slice) == 0
}

