package main

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock
// Easy

import (
    //"fmt"
    "math"
    "testing"
)

func maxProfit(prices []int) int {
    maxProfit := 0
    minValue := math.MaxInt32
    for _, v := range prices {
        if minValue > v {
            minValue = v
        }
        if diff := v - minValue; diff > maxProfit {
            maxProfit = diff
        }
    }
    return maxProfit
}

func TestMaxProfit(t *testing.T) {
    testCases := []struct {
        input []int
        expected int
    }{
        {
            input: []int{7,1,5,3,6,4},
            expected: 5,
        },
        {
            input: []int{1,101},
            expected: 100,
        },
        {
            input: []int{3, 2, 1},
            expected: 0,
        },
        {
            input: []int{0},
            expected: 0,
        },
        {
            input: []int{1},
            expected: 0,
        },
    }
    for _, testCase := range testCases {
        out := maxProfit(testCase.input)
        if out != testCase.expected {
            t.Errorf("Failed test case: %v; Out = %v", testCase, out)
        }
    }
}
