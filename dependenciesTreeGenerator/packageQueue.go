package dependenciesTreeGenerator

type queue struct {
	slice	[]*PackageNode
}

func NewQueue() queue {
	q := queue{}
	q.slice = make([]*PackageNode, 0)
	return q
}

func (q *queue) Enqueue(package_ *PackageNode){
	q.slice = append(q.slice, package_)
}

func (q *queue) Dequeue() *PackageNode {
	element := q.slice[0]
	q.slice = q.slice[1:]
	return element
}

func (q queue)IsEmpty() bool {
	return len(q.slice) == 0
}

