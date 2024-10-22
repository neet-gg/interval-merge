package merge

type Interval struct {
	Begin int
	End   int
}

// MergeInterval takes a list of intervals and merges
// all overlapping intervals.
//
// Each interval consists of two integer values
// defining its beginning and its end.
// If interval A.End equals B.Begin, A and B are considered overlapping.
// If Interval A.End + 1 equals B.Begin, A and B are considered non-overlapping
//
// Example: [{25,30}, {2,19}, {14,23}] -> [{2,23}, {25,30}]
func MergeIntervals(intervals []Interval) []Interval {
	return []Interval{}
}
