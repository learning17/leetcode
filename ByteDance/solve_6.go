package main
// https://www.nowcoder.com/practice/9ce534c8132b4e189fd3130519420cde
func solve( a []int ) int {
    size := len(a)
    if a[0] != 0 {
        return 0
    }
    if a[size-1] != size {
        return size
    }
    left, right := 0, size - 1
    for {
        if left >= right {
            return left
        }
        mid := left+(right-left)/2
        if a[mid] == mid {
            left = mid+1
        } else {
            right = mid
        }
    }
}
