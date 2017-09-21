var singleNumber = function(nums) {
    var s = nums[0]

    for (var i = 1; i < nums.length; i++) {
        s ^= nums[i]
    }

    return s
}
