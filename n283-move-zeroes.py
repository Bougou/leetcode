class Solution(object):
    def moveZeroes(self, nums):
        """
        :type nums: List[int]
        :rtype: void Do not return anything, modify nums in-place instead.
        """
        j = 0
        for i in xrange(len(nums)):
            if nums[i] != 0:
                nums[j], nums[i] = nums[i], nums[j]
                j += 1

if __name__ == '__main__':
	s = Solution()
	a = [0, 1, 0, 12, 4]
	s.moveZeroes(a)
	print a
