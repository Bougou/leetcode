class Solution(object):
    def findLUSlength(self, a, b):
        return a == b ? -1 : max(len(a), len(b))
