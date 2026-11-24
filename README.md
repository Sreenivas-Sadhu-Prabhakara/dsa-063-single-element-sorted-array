# Single Element in a Sorted Array

> Striver SDE Sheet 063 · Day 11 Binary Search — pairs start at even indexes until the single element breaks the alignment; binary search that break. Go, stdlib only, tested.

## Problem

In a sorted array where every element appears exactly twice except one, find that one — in O(log n), which rules out the XOR sweep (O(n)) this would otherwise be.

## Approach

The searchable structure is parity: left of the single element, every pair occupies indexes (even, odd); at and after it, pairs shift to (odd, even). So probe the middle, snap it to its pair's even start, and check whether the pair is intact. Intact means the single is strictly to the right (`lo = mid+2`); broken means it's here or to the left (`hi = mid`). The invariant "lo..hi always contains the single, lo always even-aligned" closes the loop with `nums[lo]` as the answer, single-element and edge positions included.

**Complexity:** O(log n) time, O(1) space.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
