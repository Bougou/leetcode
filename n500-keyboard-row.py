class Solution(object):
    def findWords(self, words):
        row1, row2, row3 = set('qwertyuiop'), set('asdfghjkl'), set('zxcvbnm')

        ret = []

        for word in words:
            w = set(word.lower())
            if w.issubset(row1) or w.issubset(row2) or w.issubset(row3):
                ret.append(word)
        
        return ret

if __name__ == '__main__':
    s = Solution()

    a = ["Hello", "Alaska", "Dad", "Peace"]
    print s.findWords(a)
