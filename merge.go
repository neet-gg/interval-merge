package merge

import "sort"

type Interval struct {
	Begin int
	End   int
}

type ByBegin []Interval

func (i ByBegin) Len() int {
	return len(i)
}

func (i ByBegin) Swap(a, b int) {
	i[a], i[b] = i[b], i[a]
}

func (i ByBegin) Less(a, b int) bool {
	return i[a].Begin < i[b].Begin
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
	// Base case
	if len(intervals) <= 1 {
		return intervals
	}

	// Sort intervals by beginning.
	// This way we only need to compare each interval to its predecessor.
	// Let I be the list of intervals:
	// Since I[i].Begin < I[i+1].Begin, if I[i].End < I[i+1].Begin
	// it follows that I[i].End < I[i+n].Begin for all n > 1.
	sort.Sort(ByBegin(intervals))

	result := []Interval{}
	prev := intervals[0]

	for i := 1; i < len(intervals); i++ {
		// intervals[i] overlaps with previous interval
		if intervals[i].Begin <= prev.End {
			// intervals[i] is not a subset of previous interval
			if intervals[i].End > prev.End {
				prev.End = intervals[i].End
			}
		} else {
			result = append(result, prev)
			prev = intervals[i]
		}
	}

	result = append(result, prev)

	return result
}
