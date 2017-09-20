class Solution(object):
    def reverseWords(self, s):
        return ' '.join([word[::-1] for word in s.split(' ')])

if __name__ == '__main__':
    s = Solution()

    a = "Let's take LeetCode contest"
    print s.reverseWords(a)
