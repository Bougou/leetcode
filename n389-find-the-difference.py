class Solution(object):
    def findTheDifference(self, s, t):
        """
        :type s: str
        :type t: str
        :rtype: str
        """
        ret = 0
        for c in s+t:
            ret ^= ord(c)
        
        return chr(ret)
            
       