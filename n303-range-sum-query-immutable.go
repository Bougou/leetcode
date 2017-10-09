package main

type NumArray struct {
    sums []int
}


func Constructor(nums []int) NumArray {
    ret := NumArray{sums: make([]int, len(nums))}
    
    for i, v := range nums {
        if i == 0 {
            ret.sums[i] = v
        } else {
            ret.sums[i] = ret.sums[i-1] + v
        }
    }
    
    return ret
}


func (this *NumArray) SumRange(i int, j int) int {
    if i == 0 {
        return  this.sums[j]
    }
    
    return this.sums[j] - this.sums[i-1]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
