class Solution(object):
    def singleNumber(self, nums):
        s = nums[0]

        for i in xrange(1, len(nums)):
            s ^= nums[i]
        
        return s
