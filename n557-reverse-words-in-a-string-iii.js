var reverseWords = function(s) {
    var ret = []

    for (var word of s.split(' ')) {
        ret.push(word.split('').reverse().join(''))
    }

    return ret.join(' ')
}

var s = "Let's take LeetCode contest"

console.log(reverseWords(s))
