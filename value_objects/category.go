package value_objects

type Category int

const (
	Book Category = iota + 1
	Seed
	Other
)
