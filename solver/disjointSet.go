package solver

type DisjointSet[T comparable] struct {
	parent map[T]T
}

func NewDisjointSet[T comparable]() DisjointSet[T] {
	return DisjointSet[T]{
		parent: make(map[T]T),
	}
}

func (d *DisjointSet[T]) Add(e ...T) {
	for _, item := range e {
		if _, exist := d.parent[item]; !exist {
			d.parent[item] = item
		}
	}
}

func (d *DisjointSet[T]) Join(e1, e2 T) {
	d.Add(e1, e2)
	p1, _ := d.Parent(e1)
	p2, _ := d.Parent(e2)

	d.parent[p1] = p2

	d.Parent(e1)
	d.Parent(e2)
}

func (d *DisjointSet[T]) Parent(e T) (T, bool) {
	var zero T
	p, exist := d.parent[e]
	if !exist {
		return zero, false
	}

	if p == e {
		return p, true
	}

	pUp, ok := d.Parent(p)
	if ok {
		d.parent[e] = pUp
		return pUp, true
	}
	return zero, false
}

func (d *DisjointSet[T]) NumberOfSets() int {
	sets := make(map[T]struct{})
	for e := range d.parent {
		if p, exist := d.Parent(e); exist {
			sets[p] = struct{}{}
		}
	}
	return len(sets)
}
