import math

class Solution(object):
    def titleToNumber(self, s):
        """
        :type s: str
        :rtype: int
        """
        
        sum = 0
        for i in xrange(len(s)):
            c = ord(s[i]) - ord('A') + 1
            exp = len(s) - i - 1
            
            sum += c * int(math.pow(26, exp))
        
        return sum

    def titleToNumber2(self, s):
        return reduce(lambda x, y: 26*x+ord(y)-64, s, 0)

if __name__ == '__main__':
    s = Solution()
    print s.titleToNumber('AAA')


