var fizzBuzz = function(n) {
    var ret = []
    for (var i = 1; i <= n; i++) {
        if (i % 15 === 0) {
            ret[i-1] = "FizzBuzz"
        } else if (i % 5 === 0) {
            ret[i-1] = "Buzz"
        } else if (i % 3 === 0) {
            ret[i-1] = "Fizz"
        } else {
            ret[i-1] = i.toString()
        }
    }
    
    return ret
};
