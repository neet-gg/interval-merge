package merge

import (
	"testing"
)

// Test two intervals A and B
// A and B do overlap and should be merged
func TestMergeOverlappingIntervals(t *testing.T) {
	merged, _ := MergeIntervals([]Interval{{10, 25}, {20, 31}})
	expected := []Interval{{10, 31}}

	if len(expected) != len(merged) {
		t.Fatalf("Expected merged intervals to be of size %d, got %d", len(expected), len(merged))
	}

	for index, interval := range merged {
		if interval != expected[index] {
			t.Errorf("Expected merged interval at index %d to be (%d, %d), got (%d, %d)", index, expected[index].Begin, expected[index].End, interval.Begin, interval.End)
		}
	}
}

// Test two intervals A and B
// A and B do not overlap and should stay separate
func TestMergeDisjointIntervals(t *testing.T) {
	input := []Interval{{10, 25}, {26, 31}}
	merged, _ := MergeIntervals(input)

	if len(input) != len(merged) {
		t.Fatalf("Expected merged intervals to be of size %d, got %d", len(input), len(merged))
	}

	for index, interval := range merged {
		if interval != input[index] {
			t.Errorf("Expected merged interval at index %d to be (%d, %d), got (%d, %d)", index, input[index].Begin, input[index].End, interval.Begin, interval.End)
		}
	}
}

// Test two intervals A and B
// A is subset of B, so only B should stay
func TestMergeSubsetIntervals(t *testing.T) {
	merged, _ := MergeIntervals([]Interval{{10, 25}, {12, 22}})
	expected := []Interval{{10, 25}}

	if len(expected) != len(merged) {
		t.Fatalf("Expected merged intervals to be of size %d, got %d", len(expected), len(merged))
	}

	for index, interval := range merged {
		if interval != expected[index] {
			t.Errorf("Expected merged interval at index %d to be (%d, %d), got (%d, %d)", index, expected[index].Begin, expected[index].End, interval.Begin, interval.End)
		}
	}
}

// Test two intervals A and B
// A is adjacent to B but does not overlap, so both should stay separate
func TestMergeAdjacentIntervals(t *testing.T) {
	input := []Interval{{10, 25}, {26, 35}}
	merged, _ := MergeIntervals(input)

	if len(input) != len(merged) {
		t.Fatalf("Expected merged intervals to be of size %d, got %d", len(input), len(merged))
	}

	for index, interval := range merged {
		if interval != input[index] {
			t.Errorf("Expected merged interval at index %d to be (%d, %d), got (%d, %d)", index, input[index].Begin, input[index].End, interval.Begin, interval.End)
		}
	}
}

// Test two intervals A and B
// A.Begin equals B.End, so they should be merged
func TestMergeTangentIntervals(t *testing.T) {
	merged, _ := MergeIntervals([]Interval{{10, 25}, {25, 35}})
	expected := []Interval{{10, 35}}

	if len(expected) != len(merged) {
		t.Fatalf("Expected merged intervals to be of size %d, got %d", len(expected), len(merged))
	}

	for index, interval := range merged {
		if interval != expected[index] {
			t.Errorf("Expected merged interval at index %d to be (%d, %d), got (%d, %d)", index, expected[index].Begin, expected[index].End, interval.Begin, interval.End)
		}
	}
}

// Test multiple intervals A, B, C, D
// All intervals are disjoint and should not be merged
func TestMergeMultipleDisjointIntervals(t *testing.T) {
	input := []Interval{{45, 47}, {99, 120}, {10, 25}, {26, 31}}
	merged, _ := MergeIntervals(input)

	if len(input) != len(merged) {
		t.Fatalf("Expected merged intervals to be of size %d, got %d", len(input), len(merged))
	}

	for index, interval := range merged {
		if interval != input[index] {
			t.Errorf("Expected merged interval at index %d to be (%d, %d), got (%d, %d)", index, input[index].Begin, input[index].End, interval.Begin, interval.End)
		}
	}
}

// Test two intervals A and B, with B having a beginning < 0
// A and B do overlap and should be merged
func TestMergeNegativeValues(t *testing.T) {
	merged, _ := MergeIntervals([]Interval{{45, 120}, {-12, 45}})
	expected := []Interval{{-12, 120}}

	if len(expected) != len(merged) {
		t.Fatalf("Expected merged intervals to be of size %d, got %d", len(expected), len(merged))
	}

	for index, interval := range merged {
		if interval != expected[index] {
			t.Errorf("Expected merged interval at index %d to be (%d, %d), got (%d, %d)", index, expected[index].Begin, expected[index].End, interval.Begin, interval.End)
		}
	}
}

// Test two intervals A and B.
// Interval B is invalid, so an error is expected
func TestMergeInvalidIntervals(t *testing.T) {
	_, err := MergeIntervals([]Interval{{12, 20}, {40, 19}})
	if err == nil {
		t.Fatalf("Expected error on invalid interval")
	}

	if err.Error() != "invalid interval" {
		t.Fatalf("Expected error %s, got %s", "invalid interval", err.Error())
	}
}
