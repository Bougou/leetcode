var addBinary = function(a, b) {
    var alen = a.length
    var blen = b.length
    var carry = 0

    var ret = ""
    for (; alen > 0 || blen > 0; ) {
        abit = 0
        bbit = 0

        if (alen > 0) {
            abit = a[alen-1] - 0
            alen--
        }

        if (blen > 0) {
            bbit = b[blen-1] - 0
            blen--
        }

        console.log(abit, bbit, carry)
        sum = abit + bbit + carry

        carry = sum >= 2 ? 1 : 0
        r = sum%2

        ret = r + ret
    }

    if (carry === 1) {
        return "1" + ret
    } else {
        return ret
    }
}
