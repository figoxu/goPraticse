package where_t

type LogicFn func() bool

func (p LogicFn) Not() Logic {
	var fn LogicFn = func() bool {
		result := p.Val()
		notResult := !result
		return notResult
	}
	return fn
}

func (p LogicFn) Or(vs ...Logic) Logic {
	var fn LogicFn = func() bool {
		fns := append(vs, p)
		for _, f := range fns {
			v := f.Val()
			if v {
				return true
			}
		}
		return false
	}
	return fn
}

func (p LogicFn) And(vs ...Logic) Logic {
	var fn LogicFn = func() bool {
		fns := append(vs, p)
		for _, f := range fns {
			v := f.Val()
			if !v {
				return false
			}
		}
		return true
	}
	return fn
}

func (p LogicFn) Val() bool {
	return p()
}
