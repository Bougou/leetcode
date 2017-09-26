/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number[]}
 */
var intersection = function(nums1, nums2) {
    var ret = []
    
    for (let n1 of nums1) {
        if (nums2.includes(n1) && ! ret.includes(n1)) {
            ret.push(n1)
        }
    }
    
    return ret
};
