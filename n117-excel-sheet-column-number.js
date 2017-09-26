var titleToNumber = function(s) {
    var sum = 0
    for (var i = 0; i < s.length; i++) {
        var k = s[i].charCodeAt(0) - 'A'.charCodeAt(0) + 1
        var exp = s.length - i - 1
        sum += k * Math.pow(26, exp)
    }
    
    return sum
};
