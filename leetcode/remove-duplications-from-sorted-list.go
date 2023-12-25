package main

/**
 * https://leetcode.com/problems/remove-duplicates-from-sorted-list/description/
 * Easy
 */

import (
    //"fmt"
    "testing"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    res := head
    for head != nil && head.Next != nil {
        if head.Val == head.Next.Val {
            head.Next = head.Next.Next
        } else {
            head = head.Next
        }
    }
    return res
}

