var findDisappearedNumbers = function(nums) {
    nums.forEach((value, index) => {
        newIndex = Math.abs(value) - 1
        nums[newIndex] = -Math.abs(nums[newIndex])
    })
    
    console.log(nums)

    var ret = []
    nums.forEach((value, index) => {
        if (value > 0) {
            ret.push(index + 1)
        }
    })

    return ret
}


var a = [4,3,2,7,8,2,3,1]

console.log(findDisappearedNumbers(a))
