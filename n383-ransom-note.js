var canConstruct = function(ransomNote, magazine) {
    let a = new Array(26).fill(0)
    for (let i = 0; i < magazine.length; i++) {
        let index = magazine[i].charCodeAt(0) - 'a'.charCodeAt(0)
        a[index]++
    }
    
    for (let j = 0; j < ransomNote.length; j++) {
        let index = ransomNote[j].charCodeAt(0) - 'a'.charCodeAt(0)
        if (--a[index] < 0) {
            return false
        }
    }
    
    return true
};
