package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Device struct {
	pattern [][]rune
	pins    []int
}

func (d Device) Verify() {
	for _, runes := range d.pattern {
		for _, r := range runes {
			fmt.Printf("%c", r)
		}
		fmt.Println()
	}
	fmt.Println(d.pins)
}

func readInput(path string) (locks, keys []Device, err error) {
	// Open the input file
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, nil, err
	}
	defer file.Close()

	// Read and parse the file contents
	scanner := bufio.NewScanner(file)
	var block [][]rune

	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" { // Empty line separates blocks
			if len(block) > 0 {
				isLock := isFullHashAtTop(block)
				if isLock {
					locks = append(locks, Device{
						pattern: block,
						pins:    calculatePinHeights(block),
					})
				} else {
					keys = append(keys, Device{
						pattern: block,
						pins:    calculatePinHeights(block),
					})
				}
				block = nil
			}
		} else {
			block = append(block, []rune(line))
		}
	}

	// Append the final block if existed
	if len(block) > 0 {
		isLock := isFullHashAtTop(block)
		if isLock {
			locks = append(locks, Device{
				pattern: block,
				pins:    calculatePinHeights(block),
			})
		} else {
			keys = append(keys, Device{
				pattern: block,
				pins:    calculatePinHeights(block),
			})
		}
	}

	if err = scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil, nil, err
	}

	return keys, locks, err
}
