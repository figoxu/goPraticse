package think

type TestError struct {
	Value string
}

func (e TestError) Error() string {
	return e.Value
}
