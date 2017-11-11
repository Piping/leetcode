package main

import "fmt"

type CpSorter interface {
	split(ar []int, low uint, mid *uint, high uint)
	join(ar []int, low, mid, high uint)
}

type MergeSorter struct{}

func (s MergeSorter) split(ar []int, low uint, mid *uint, high uint) {
	*mid = (high-low+1)/2 + low
}
func (s MergeSorter) join(ar []int, low, mid, high uint) {
	lo, hi := low, mid
	temp := make([]int, len(ar))
	for k := uint(0); k <= high-low; k++ {
		if lo >= mid {
			temp[k] = ar[hi]
			hi++
		} else if hi > high {
			temp[k] = ar[lo]
			lo++
		} else {
			if ar[lo] > ar[hi] {
				temp[k] = ar[hi]
				hi++
			} else {
				temp[k] = ar[lo]
				lo++
			}
		}
	}
	for k := uint(0); k <= high-low; k++ {
		ar[k+low] = temp[k]
	}
}

func sort(s CpSorter, ar []int, low, high uint) {
	if high >= uint(len(ar)) {
		return
	}
	if low < high {
		var mid uint
		s.split(ar, low, &mid, high)
		sort(s, ar, low, mid-1)
		sort(s, ar, mid, high)
		s.join(ar, low, mid, high)
	}
}

func main() {
	fmt.Println("hi")
	ar := []int{1, 3, 2, 5, 9, 7, 8}
	sort(MergeSorter{}, ar, 0, uint(len(ar)-1))
	fmt.Println(ar)
}
