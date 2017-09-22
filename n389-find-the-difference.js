var findTheDifference = function(s, t) {
    var ret = 0
    var data = s + t
    data.split('').forEach(c => {
        ret ^= c.charCodeAt(0)
    })
    
    return String.fromCharCode(ret)
};
