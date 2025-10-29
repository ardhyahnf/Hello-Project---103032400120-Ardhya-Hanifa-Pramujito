package main

import "fmt"

func BinSearch2(tab tabInt, n, x int)int{
	var left, right, mid int
	var idx int
	
	left = 0
	right = n-1
	idx = -1 
	for left <= right && idx != -1 {
		mid = (left + right)/2
		if x < tab[mid]{
			right = mid - 1 
		} else if x > tab[mid] {
			left = mid + 1
		} else {
			idx = mid
		}
	}
	return idx
}