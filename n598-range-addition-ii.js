var maxCount = function(m, n, ops) {
    if (ops.length === 0) {
        return m*n
    }
    
    let x_arr = []
    let y_arr = []
    
    ops.forEach(ele => {
        x_arr.push(ele[0])
        y_arr.push(ele[1])
    })
    
    min_x = Math.min(...x_arr)
    min_y = Math.min(...y_arr)
    
    return min_x * min_y
};


var maxCount2 = function(m, n, ops) {
    let min_x = m
    let min_y = n
    
    for (let i = 0; i < ops.length; i++) {
        if (ops[i][0] < min_x) min_x = ops[i][0];
        if (ops[i][1] < min_y) min_y = ops[i][1];
    }
    
    return min_x * min_y
};
