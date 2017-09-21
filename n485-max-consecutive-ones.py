class Solution(object):
    def findMaxConsecutiveOnes(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        maxv = 0
        cons = 0
        
        for i in xrange(len(nums)):
            if nums[i] == 1:
                cons += 1
                maxv = max(maxv, cons)
            else:
                cons = 0
        
        return maxv
