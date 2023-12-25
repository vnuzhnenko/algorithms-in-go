package main

/**
 * https://leetcode.com/problems/merge-two-sorted-lists/submissions/1126516807/
 * Easy
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    return doMergeTwoLists(list1, list2)
}

func doMergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    if list1 == nil && list2 == nil {
        return nil
    }

    var current *ListNode
    if list1 == nil {
        current = list2
        current.Next = doMergeTwoLists(nil, list2.Next)
    } else if list2 == nil {
        current = list1
        current.Next = doMergeTwoLists(list1.Next, nil)
    } else if list1.Val < list2.Val {
        current = list1
        current.Next = doMergeTwoLists(list1.Next, list2)
    } else {
        current = list2
        current.Next = doMergeTwoLists(list1, list2.Next)
    }
    return current
}

// See a very clever solution: https://leetcode.com/problems/merge-two-sorted-lists/solutions/360472/golang-0ms-recursive-solution/?source=submission-ac
