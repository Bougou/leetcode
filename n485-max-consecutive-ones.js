var findMaxConsecutiveOnes = function(nums) {
    var max = 0
    var cons = 0
    
    for (var i = 0; i < nums.length; i++) {
        if (nums[i] === 1) {
            cons += 1
            max = max > cons ? max : cons
        } else {
            cons = 0
        }
    }
    
    return max
};
