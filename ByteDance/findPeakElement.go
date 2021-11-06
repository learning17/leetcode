package main
//https://www.nowcoder.com/practice/fcf87540c4f347bcb4cf720b5b350c76
func findPeakElement( nums []int ) int {
    left, right := 0, len(nums)-1
    if right == 0 { 
        return 0
    }   
    if nums[left] > nums[left+1] {
        return left
    }   
    if nums[right] > nums[right-1] {
        return right
    }   
    for left < right {
        mid := left + (right - left)/2
        if nums[mid] > nums[mid+1] && nums[mid] > nums[mid-1] {
            return mid 
        } else if nums[mid+1] > nums[mid] {                                                                          
            left = mid+1
        } else {
            right = mid-1
        }   
    }   
    if left == right {
        return left
    }   
    return -1  
}
