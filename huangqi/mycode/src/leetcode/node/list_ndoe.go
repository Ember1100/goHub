package node

/**
 * 反正链表
 */
func ReverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var next, pre *ListNode
	for head != nil {
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}

//328. 奇偶链表
func OddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var odd *ListNode = head
	var even *ListNode = head.Next
	var evenHead *ListNode = even
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next

		even.Next = odd.Next
		even = even.Next
	}

	odd.Next = evenHead
	return head
}

// 83. 删除排序链表中的重复元素
func DeleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	fast := head.Next // 快指针，用于遍历链表
	slow := head      // 慢指针，指向当前不重复的节点

	for fast != nil {
		if fast.Val != slow.Val {
			// 当快指针的值不等于慢指针的值时，表示找到了一个新的不重复节点
			slow.Next = fast // 将新的不重复节点连接到慢指针的后面
			slow = slow.Next // 移动慢指针
		} else {
			slow.Next = nil // 当快指针的值等于慢指针的值时，表示出现了重复节点，将慢指针的Next置为nil，断开重复节点的连接
		}
		fast = fast.Next // 移动快指针
	}

	return head // 返回去重后的链表
}

//82. 删除排序链表中的重复元素 II
func DeleteDuplicatesII(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var listNode ListNode = ListNode{Val: 0}
	dummy := &listNode
	temp := dummy
	current := head
	for current != nil {
		if current.Next != nil && current.Val == current.Next.Val {
			for current.Next != nil && current.Val == current.Next.Val {
				current = current.Next
			}
			current = current.Next
		} else {
			temp.Next = current
			temp = temp.Next
			current = current.Next
			temp.Next = nil
		}
	}
	return dummy.Next
}
