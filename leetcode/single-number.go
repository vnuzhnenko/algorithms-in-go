package main

/**
 * https://leetcode.com/problems/single-number/description/
 * Easy
 */

func singleNumber(nums []int) int {
    return sortedSolution(nums)
}

/**
 * O(NLogN) + O(N) 
 */
func sortedSolution(nums []int) int {
    slices.Sort(nums)
    for i := 0; i < len(nums); i += 2 {
        if i == len(nums) - 1 || nums[i] != nums[i + 1] {
            return nums[i]
        }
    }
    return -1
}

// Beautiful solution with XOR because this operation is commutative and associative
// Intuition:
// 1 xor 2 xor 3 xor 1 xor 2 xor 3 xor 4 = (1 xor 1) xor (2 xor 2) xor (3 xor 3) xor 4 (commutative & associative)
// = 0 xor 0 xor 0 xor 4
// = 4
func xorSolution(nums []int) int {
    result := 0
    for _, num := range nums {
        result = num ^ result
    }
    return result
}

