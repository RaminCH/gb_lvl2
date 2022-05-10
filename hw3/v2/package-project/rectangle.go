package rectangle

type Rectangle struct {
	breadth int
	len     int
}

func (r *Rectangle) Area() int {
	return r.len * r.breadth
}
