package main
// https://www.nowcoder.com/practice/c56f6c70fb3f4849bc56e33ff2a50b6b

type ListNode struct{
	Val int
	Next *ListNode
}

func addInList( head1 *ListNode ,  head2 *ListNode ) *ListNode {
	stack1, stack2 := []int{}, []int{}
	for ;head1 != nil; head1 = head1.Next {
		stack1 = append(stack1, head1.Val)
	}
	for ;head2 != nil; head2 = head2.Next {
		stack2 = append(stack2, head2.Val)
	}
	arr := []int{}
	bit := 0
	for len(stack1) > 0 && len(stack2) > 0 {
		num1, num2 := stack1[len(stack1)-1], stack2[len(stack2)-1]
		stack1, stack2 = stack1[:len(stack1)-1], stack2[:len(stack2)-1]
		ans := num1 + num2 + bit
		bit = ans / 10
		arr = append(arr, ans%10)
	}
	for len(stack1) > 0 {
		num1 := stack1[len(stack1)-1]
		stack1 = stack1[:len(stack1)-1]
		ans := num1 + bit
		bit = ans / 10
		arr = append(arr, ans%10)
	}
	for len(stack2) > 0 {
		num2 := stack2[len(stack2)-1]
		stack2 = stack2[:len(stack2)-1]
		ans := num2 + bit
		bit = ans / 10
		arr = append(arr, ans%10)
	}
	if bit != 0 {
		arr = append(arr, bit)
	}
	root := &ListNode{}
	node := root
	for i := len(arr)-1; i >= 0; i-- {
		node.Next = &ListNode{
			Val: arr[i],
		}
		node = node.Next
	}
	return root.Next
}
