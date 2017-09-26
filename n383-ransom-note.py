class Solution(object):
    def canConstruct(self, ransomNote, magazine):
        """
        :type ransomNote: str
        :type magazine: str
        :rtype: bool
        """
        mag = [0] * 26
        
        for c in magazine:
            index = ord(c) - ord('a')
            mag[index] += 1
        
        for c in ransomNote:
            index = ord(c) - ord('a')
            mag[index] -= 1
            if mag[index] < 0:
                return False
        
        return True
