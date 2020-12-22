package common

type Constraint struct {
	Inter1 Interval
	Inter2 Interval
}

func (c *Constraint) Satisfies(n int) bool {
	return (c.Inter1.Min <= n && n <= c.Inter1.Max) || (c.Inter2.Min <= n && n <= c.Inter2.Max)
}

// Interval is a closed interval [Min,Max]
type Interval struct {
	Min int
	Max int
}

// IsIn returns true if Min <= n <= Max of the receiver.
func (i *Interval) IsIn(n int) bool {
	return i.Min <= n && n <= i.Max
}
