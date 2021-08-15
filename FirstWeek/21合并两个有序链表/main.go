package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{Val: 0}
	cur := result
	//先创建一个前置节点，来确定它下一个next指向谁
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			//目标是合并为升序链表所以小的值得在前面
			cur.Next = l1
			l1 = l1.Next

		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	//在两个链表同时不为nil的情况下遍历完成后，会存在有l1 or l2还存在的情况下
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	//因为前置节点使我们自己创建的，所以返回的时候需要返回.Next
	return result.Next

}

func main() {
	mergeTwoLists()

}
