/**
 * @param {string} word
 * @return {boolean}
 */
var detectCapitalUse = function(word) {
    var capCount = 0
    for (var i = 0; i < word.length; i++) {
        if ('Z'.charCodeAt(0) - word[i].charCodeAt(0) >= 0) {
            capCount++
        }
    }
    
    if (capCount === word.length || 
        capCount === 0 ||
        (capCount === 1 && 'Z'.charCodeAt(0) - word[0].charCodeAt(0) >= 0)
    ) {
        return true
    }
    
    return false
};
