# interval-merge

## Specification

Given a list of intervals $I = \{i_1,\dots,i_n\}$ consisting of an interger value $b_j$ for the beginning of the interval and an integer value $e_j$ representing the end of the interval, return a list of intervals $I_m$ so that all overlapping intervals in $I$ are merged together.
Intervals $i_j$ and $i_k$ are considered "overlapping", if $e_j \geq b_k$ .

### Example

**Input:** [25,30] [2,19] [14,23]

**Output:** [2,23] [25,30] 

## Analysis

### Correctness

If all intervals are overlapping intervals, it is obvious that comparing only the current and the previous interval 
is sufficient. Therefore we will only consider the case of non-overlapping intervals. In that case, it is 
necessary to compare the current interval to all previous intervals since they are all separate.
By sorting the intervals first by their beginning, we can avoid this since if interval $i_j$ and $i_{j+1}$ do not overlap then interval $i_j$ does also not overlap with all following intervals $i_{j+n}$ for $1 < n < \mid I\mid - j$.
Let $b_j$ be the beginning of the $j$-th interval and $e_j$ the end of the $j$-th interval.
If $b_{j} \leq b_{j+1}$ and $e_{j} < b_{j+1}$ (non-overlapping), then it follows that there cannot be any $b_{j+n} < e_j$ for $n>1$ which is also $> b_{j+1}$.

### Runtime

The given function sorts the input list using `go sort` in $O(n\log n)$. [^1]
After that only $n-1$ compares are neeeded, where each compare can be performed in $O(1)$.
Appending is also in (amortized) $O(1)$.
Therefore the total runtime of `MergeIntervals` is in $O(n\log n)$.

### Memory

No additional memory is required beyond the result list, which is in $O(n)$.


[^1]:   https://pkg.go.dev/sort