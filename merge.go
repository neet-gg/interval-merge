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
// Each interval consists of two integer values defining its beginning and its end.
// If interval A.End is greater or equal to B.Begin, A and B are considered overlapping.
func MergeIntervals(intervals []Interval) []Interval {
	if len(intervals) <= 1 {
		return intervals
	}

	// Sort intervals by beginning.
	// This way we only need to compare each interval to its predecessor.
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
