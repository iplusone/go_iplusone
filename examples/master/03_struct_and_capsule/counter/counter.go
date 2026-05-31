package counter

// Counter は非公開フィールド n を持つ。外から直接書き換えられない。
type Counter struct {
	n int
}

func (c *Counter) Increment() {
	c.n++
}

func (c *Counter) Reset() {
	c.n = 0
}

func (c *Counter) Value() int {
	return c.n
}
