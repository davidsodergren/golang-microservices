package utils

import "sort"

func BubbleSort(elements []int) {
	keepRunnig := true
	
	for keepRunnig {
		keepRunnig = false

		for i := 0; i < len(elements) -1; i++ {
			if elements[i] > elements[i+1] {
				elements[i], elements[i+1] = elements[i+1], elements[i]
				keepRunnig = true
			}
		}
	}
}

func Sort(elements []int)  {
	if len(elements) > 1000 {
		sort.Ints(elements)
	} else {
		BubbleSort(elements)
	}
}
