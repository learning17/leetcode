package main
// 最大的异或 https://leetcode-cn.com/problems/ms70jA/

func findMaximumXOR(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
    t := Constructor()
    for _, num := range nums {
        t.insert(num)
    }
    maxXOR := nums[0] ^ nums[1]
    for _, num := range nums {
        ans := t.searchMaxXOR(num)
        if ans > maxXOR {
            maxXOR = ans
        }
    }
    return maxXOR
}

type TrieNode struct {
    children [2]*TrieNode
}

func Constructor() *TrieNode{
    return &TrieNode{}
}

func (t *TrieNode) insert(value int) {
    cur := t
    for i := 30; i >= 0; i-- {
        bit := (value >> i) & 1
        if cur.children[bit] == nil {
            cur.children[bit] = &TrieNode{}
        }
        cur = cur.children[bit]
    }
}

func (t *TrieNode) searchMaxXOR(value int) int {
    ans, cur := 0, t
    for i := 30; i >= 0; i-- {
        bit := (value >> i) & 1
        bit ^= 1
        if cur.children[bit] == nil {
            bit ^= 1
        }
        cur = cur.children[bit]
        ans |= (bit << i)
    }
    return value ^ ans
}
