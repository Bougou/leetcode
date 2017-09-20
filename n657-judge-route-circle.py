class Solution(object):
    def judgeCircle(self, moves):
        """
        :type moves: str
        :rtype: bool
        """
        
        path = {"R": [1, 0], "L": [-1, 0], "U": [0, 1], "D": [0, -1]}
        
        x, y = 0, 0
        
        for m in moves:
            x += path[m][0]
            y += path[m][1]
        
        return x == 0 and y == 0

if __name__ == '__main__':
    s = Solution()

    print s.judgeCircle("RULD")
    print s.judgeCircle("RULR")
