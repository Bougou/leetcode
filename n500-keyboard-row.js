var findWords2 = function(words) {
    let row1 = new Set('qwertyuiop')
    let row2 = new Set('asdfghjkl')
    let row3 = new Set('zxcvbnm')

    let ret = []

    for (let word of words) {
        let flag = true
        let lword = word.toLowerCase()
        let first = lword[0]


        let row
        if (row1.has(first)) {
            row = row1
        } else if (row2.has(first)) {
            row = row2
        } else {
            row = row3
        }


        wordset = new Set(lword)

        for (let w of wordset) {
            if (!row.has(w)) {
                flag = false
            }
        }

        if (flag) {
            ret.push(word)
        }
    }

    return ret
}

var findWords = function(words) {
    var row1exp = /^[qwertyuiop]*$/i
    var row2exp = /^[asdfghjkl]*$/i
    var row3exp = /^[zxcvbnm]*$/i

    return words.filter(word => {
        return row1exp.test(word) || row2exp.test(word) || row3exp.test(word) || false
    })
}

a = ["Hello", "Alaska", "Dad", "Peace"]

console.log(findWords(a))
console.log(findWords2(a))
