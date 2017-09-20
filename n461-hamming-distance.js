function hammingDistance(x, y) {
    let xor = x ^ y

    console.log(x.toString(2))
    console.log(y.toString(2))
    console.log(xor.toString(2))
    
    let count = 0

    for (let i = 0; i < 32; i++) {
        count += (xor >> i) & 1
    }
    return count
}

console.log(hammingDistance(18,49))
