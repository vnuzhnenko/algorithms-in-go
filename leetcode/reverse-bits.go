package main

/**
 * https://leetcode.com/problems/reverse-bits/description/
 * Easy
 */

func reverseBits(num uint32) uint32 {
    res := num & 1
    for i := 1; i < 32; i++ {
        res = res << 1
        res = res | ((num >> i )& 1)
    }
    return res
}
