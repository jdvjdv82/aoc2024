package main

import (
	"fmt"
	"sort"
)

func main() {
	locks, keys, err := readInput("inputs/25")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Sort locks by their pin heights
	sort.Slice(locks, func(i, j int) bool {
		return comparePinHeights(locks[i].pins, locks[j].pins)
	})

	// Sort keys by their pin heights
	sort.Slice(keys, func(i, j int) bool {
		return comparePinHeights(keys[i].pins, keys[j].pins)
	})

	fmt.Println("\nMatching Keys to Locks:")

	totalsFits := 0

	lockMatched := make(map[int]struct{})
	keyMatched := make(map[int]struct{}) // Track which keys have been matched

	for kIdx, key := range keys {
		matched := make([]int, 0, len(locks))
		for lIdx, lock := range locks {
			if fits(key.pins, lock.pins) {
				matched = append(matched, lIdx+1)
				lockMatched[lIdx] = struct{}{} // Mark this lock as matched
				keyMatched[kIdx] = struct{}{}  // Mark this key as matched
				totalsFits++
			}
		}

		if len(matched) == 0 {
			fmt.Printf("Key %d does not fit any lock.\n", kIdx+1)
		} else {
			fmt.Printf("Key %d fits Locks %v\n", kIdx+1, matched)
		}
	}

	// Print summary
	fmt.Printf("\nResult:\n")
	fmt.Printf("Number of matched locks: %d out of %d\n", len(lockMatched), len(locks))
	fmt.Printf("Number of matched keys: %d out of %d\n", len(keyMatched), len(keys))
	fmt.Printf("Total matched key/lock combinations: %d\n", totalsFits)
}

// Checks if the top row of the block is a full row of hashes ('#')
func isFullHashAtTop(block [][]rune) bool {
	if len(block) == 0 {
		return false
	}
	for _, r := range block[0] {
		if r != '#' {
			return false
		}
	}
	return true
}

func calculatePinHeights(block [][]rune) []int {
	cols := len(block[0]) // Assumes all rows are the same length
	pinHeights := make([]int, cols)

	for col := 0; col < cols; col++ {
		height := 0
		for row := len(block) - 1; row >= 0; row-- { // Start from bottom
			if block[row][col] == '#' {
				height++
			}
		}
		pinHeights[col] = height - 1
	}

	return pinHeights
}

// Helper function to compare pin heights lexicographically
func comparePinHeights(a, b []int) bool {
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			return a[i] < b[i]
		}
	}
	return len(a) < len(b)
}

const maxPinHeight = 5

// Helper function to check if a key fits a lock
func fits(key, lock []int) bool {
	if len(key) != len(lock) {
		return false
	}
	for i := 0; i < len(key); i++ {
		if key[i]+lock[i] > maxPinHeight { // A key pin must be shorter or equal to the lock pin
			return false
		}
	}
	return true
}
