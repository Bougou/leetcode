package main

func plusOne(digits []int) []int {
    carry := 1
    res := []int{}
    
    for i := len(digits)-1; i >= 0; i-- {
        sum := digits[i] + carry
        if (sum == 10 ) {
            carry = 1
            res = append([]int{0}, res...)
        } else {
            carry = 0
            res = append([]int{sum}, res...)
        }         
    }
    
    if carry == 1 {
        res = append([]int{1}, res...)
    }
    
    return res
}
