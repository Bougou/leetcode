class Solution(object):
    def arrayPairSum(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        nums.sort()
        print nums

        # print sorted(nums)
        
        sum = 0
        for i in xrange(0, len(nums), 2):
            sum += nums[i]
            
        return sum

if __name__ == '__main__':
    a = [1, 3, 2, 4]
    
    s = Solution()
    print s.arrayPairSum(a) 
