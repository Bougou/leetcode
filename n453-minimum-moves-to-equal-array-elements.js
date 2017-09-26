/**
 * @param {number[]} nums
 * @return {number}
 */
var minMoves = function(nums) {
    return nums.reduce((i, j) => i+j, 0) - Math.min(...nums) * nums.length
};
