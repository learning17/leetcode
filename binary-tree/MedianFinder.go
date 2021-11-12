package main
// https://leetcode-cn.com/problems/find-median-from-data-stream/

type MedianFinder struct {
	arr []int
	size int
}

func Constructor() MedianFinder {
	return MedianFinder{}
}

func (this *MedianFinder) AddNum(num int)  {
	if this.size == 0 {
		this.arr = append(this.arr, num)
		this.size++
		return
	}
	left, right := 0, this.size-1
	for left <= right {
		mid := left + (right-left)/2
		if this.arr[mid] <= num {
			left = mid+1
		} else {
			right = mid - 1
		}
	}
	this.arr = append(this.arr[:left], append([]int{num}, this.arr[left:]...)...)
	this.size++
}

func (this *MedianFinder) FindMedian() float64 {
	if this.size % 2 == 0 {
		return (float64(this.arr[this.size/2]) + float64(this.arr[this.size/2-1]))/2
	} else {
		return float64(this.arr[this.size/2])
	} 
}
