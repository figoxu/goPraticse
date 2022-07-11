package where_t

type WhereT[T any] func(x *T) bool

func (p WhereT[T]) Not() WhereT[T] {
	var fn WhereT[T] = func(x *T) bool {
		ok := p(x)
		return !ok
	}
	return fn
}

func (p WhereT[T]) Or(vs ...WhereT[T]) WhereT[T] {
	var fn WhereT[T] = func(x *T) bool {
		fns := append(vs, p)
		for _, f := range fns {
			v := f(x)
			if v {
				return true
			}
		}
		return false
	}
	return fn
}

func (p WhereT[T]) And(vs ...WhereT[T]) WhereT[T] {
	var fn WhereT[T] = func(x *T) bool {
		fns := append(vs, p)
		for _, f := range fns {
			v := f(x)
			if !v {
				return false
			}
		}
		return true
	}
	return fn
}

func (p WhereT[T]) Val(vs ...*T) []*T {
	var out []*T
	for _, v := range vs {
		if p(v) {
			out = append(out, v)
		}
	}
	return out
}
