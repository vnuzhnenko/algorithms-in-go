package main

/**
 * https://leetcode.com/problems/number-complement/description/
 * Easy
 */

func findComplement(num int) int {
    return (1 << bits.Len(uint(num)) - 1) ^ num
}
