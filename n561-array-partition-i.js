var arrayPairSum = function(nums) {
    nums.sort((a, b) => a - b)
    console.log(nums)
    
    var sum = 0
    for (var i = 0; i < nums.length; i += 2) {
        sum += nums[i]
    }

    return sum
}

let a = [1, 3, 2, 4]
console.log(arrayPairSum(a))
