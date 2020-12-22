package common

// Interval is a closed interval [Min,Max]
type Interval struct {
	Min int
	Max int
}

// IsIn returns true if Min <= n <= Max of the receiver.
func (i *Interval) IsIn(n int) bool {
	return i.Min <= n && n <= i.Max
}
