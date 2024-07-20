package util

import (
	"fmt"
	"strconv"
	"strings"
)

type Ranges []*Range

type Range struct {
	start int
	end   int
}

func (r Range) Contain(index int) bool {
	return index <= r.end && index >= r.start
}

func (rs Ranges) Contain(index int) bool {
	for _, r := range rs {
		if r.Contain(index) {
			return true
		}
	}
	return false
}

func ParseRanges(rgs string) (Ranges, error) {
	var ranges []*Range
	if rgs == "" {
		return nil, fmt.Errorf("empty ranges")
	}
	pairs := strings.Split(rgs, ",")
	for _, p := range pairs {
		r, err := ParseRange(p)
		if err != nil {
			return nil, err
		}
		ranges = append(ranges, r)
	}
	return ranges, nil
}

func ParseRange(rg string) (*Range, error) {
	pairs := strings.SplitN(rg, "-", 2)
	if len(pairs) != 2 {
		pairs = append(pairs, pairs[0])
	}
	start, err := strconv.Atoi(pairs[0])
	if err != nil {
		return nil, fmt.Errorf("invalid range '%s', expect a number for start", rg)
	}
	end, err := strconv.Atoi(pairs[1])
	if err != nil {
		return nil, fmt.Errorf("invalid range '%s', expect a number for end", rg)
	}
	return &Range{
		start: start,
		end:   end,
	}, nil
}
