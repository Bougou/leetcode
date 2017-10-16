# The isBadVersion API is already defined for you.
# @param version, an integer
# @return a bool
# def isBadVersion(version):

class Solution(object):
    def firstBadVersion(self, n):
        """
        :type n: int
        :rtype: int
        """
        s = 1
        e = n
        while s < e:
            mid = (s + e) / 2
            if not isBadVersion(mid):
                s = mid + 1
            else:
                e = mid
        
        return s
