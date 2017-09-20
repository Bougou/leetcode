class Solution(object):
    def hammingDistance(self, x, y):
        xor = x ^ y
        print "%s\n%s\n%s\n" % (bin(x), bin(y), bin(xor))
        count = 0
        for i in xrange(32):
            count += (xor >> i) & 1
        return count


if __name__ == '__main__':
    s = Solution()

    r = s.hammingDistance(18,49)
    print r
