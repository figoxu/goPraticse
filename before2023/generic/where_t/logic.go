package where_t

type Logic interface {
	Not() Logic
	Or(vs ...Logic) Logic
	And(vs ...Logic) Logic
	Val() bool
}
