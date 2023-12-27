package main

/**
 * https://leetcode.com/problems/two-sum/description/
 * Easy
 */
import (
    "testing"
    //"fmt"
)

func twoSum(nums []int, target int) []int {
    return hashBased(nums, target)
}

// O(2n)
func hashBased(nums []int, target int) []int {
    hash := make(map[int]int, len(nums))
    for idx, num := range nums {
        hash[target - num] = idx
    }
    for idx1, num := range nums {
        idx2, found := hash[num]
        if found && idx1 != idx2 {
            return []int{idx1, idx2}
        }
    }
    return []int{0, 0}
}

// O(n^2)
func bruteForce(nums []int, target int) []int {
    for i := 0; i < len(nums) - 1; i++ {
        for j := i + 1; j < len(nums); j++ {
           if nums[i] + nums[j] == target {
               return []int{i, j}
           }
        }
    }
    return []int{}
}

func TestTwoSum(t *testing.T) {
    testCases := []struct{
        input []int
        target int
        expected []int
    } {
        {
            input: []int{3,3},
            target: 6,
            expected: []int{0,1},
        },
        {
            input: []int{2,7,11,15},
            target: 9,
            expected: []int{0,1},
        },
        {
            input: []int{3,2,4},
            target: 6,
            expected: []int{1,2},
        },
    }
    for _, testCase := range testCases {
        res := hashBased(testCase.input, testCase.target)
        if res[0] != testCase.expected[0] || res[1] != testCase.expected[1] {
            t.Errorf("got value: %v; Expected: %v", res, testCase.expected)
        }
    }
}
