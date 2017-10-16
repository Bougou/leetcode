/**
 * @param {number[]} nums
 * @return {string}
 */
var largestNumber = function(nums) {
    nums.sort((a, b) => (b + '' + a) - (a + '' + b))
    
    ret = nums.join('')
    
    if (ret - 0 === 0) {
        return '0'
    }

    return ret
};
