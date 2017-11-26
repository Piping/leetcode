package main

import "fmt"

type CpSorter interface {
	split(ar []int, low uint, mid *uint, high uint)
	join(ar []int, low, mid, high uint)
}
type Comparable interface {
	Less(i, j uint) bool
	Greater(i, j uint) bool
	Equal(i, j uint) bool
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

type QuickSorter struct{}

func (s QuickSorter) split(ar []int, low uint, mid *uint, high uint) {
	key := ar[high]
	i := low
	for j := low; j <= high; j++ {
		if ar[j] <= key {
			ar[i], ar[j] = ar[j], ar[i]
			i++
		}
	}
	if i > high {
		*mid = high
	} else {
		*mid = i
	}
}
func (s QuickSorter) join(ar []int, low, mid, high uint) {}

type SelectionSorter struct{}

func (s SelectionSorter) split(ar []int, low uint, mid *uint, high uint) {
	max := low
	for i := low + 1; i <= high; i++ {
		if ar[max] < ar[i] {
			max = i
		}
	}
	ar[max], ar[high] = ar[high], ar[max]
	*mid = high
}

func (s SelectionSorter) join(ar []int, low, mid, high uint) {}

type InsertionSorter struct{}

func (s InsertionSorter) split(ar []int, low uint, mid *uint, high uint) {
	*mid = high
}
func (s InsertionSorter) join(ar []int, low, mid, high uint) {
	key, j := ar[mid], mid-1
	//!unsigned int overflow is positive, uint >= 0
	for j > low && ar[j] > key {
		ar[j+1] = ar[j]
		j--
	}
	if j == low {
		ar[j+1] = ar[j]
		j--
	}
	ar[j+1] = key
}
func sort(s CpSorter, ar []int, low, high uint) {
	if high >= uint(len(ar)) {
		return
	}
	if low < high {
		var mid uint
		//precondition low < high, low and high differ at least by one
		//postcondition low < mid <= high
		s.split(ar, low, &mid, high)
		sort(s, ar, low, mid-1)
		sort(s, ar, mid, high)
		s.join(ar, low, mid, high)
	}
}

type test struct {
	num int
}

//! golang pass by value
//! receiver call can be invokeed on nil, but no deference
func (t *test) up() *test {
	if t == nil {
		return new(test)
	}
	return t
}
func main() {
	fmt.Println("hello world")
	var t *test
	t = t.up()
	fmt.Println(t)

	ar := []int{12, 13, 14, 15, 4, 3, 2, 1, 4, 5, 6, 9, 8, 7}
	sort(InsertionSorter{}, ar, 0, uint(len(ar)-1))
	fmt.Println(ar)
}
