package domain

import (
	"calc-feed-nums/app/transport/dto"
	"sort"
)

type Numbers struct {
	Numbers map[int]struct{}
}

// Constructor
func NewNumbers() Numbers {
	var n Numbers
	n.Numbers = make(map[int]struct{})
	return n
}

// Add new numbers from feed
func (numbers *Numbers) Add(feedNums dto.Numbers) {
	for _, num := range feedNums.Numbers {
		numbers.Numbers[num] = struct{}{}
	}
}

// Convert to transfer object for returning the result
func (numbers Numbers) ConvertToTransfer() dto.Numbers {
	keys := make([]int, 0)
	for k := range numbers.Numbers {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return dto.Numbers{
		Numbers: keys,
	}
}
