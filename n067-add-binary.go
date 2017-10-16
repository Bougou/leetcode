func addBinary(a string, b string) string {
	aLen := len(a)
	bLen := len(b)
	// 进位
	carry := 0
	var ret string

	for aLen > 0 || bLen > 0 {
		aBit := 0
		bBit := 0

		if aLen > 0 {
			aBit = int(a[aLen-1] - '0')
			aLen--
		}

		if bLen > 0 {
			bBit = int(b[bLen-1] - '0')
			bLen--
		}

		sum := aBit + bBit + carry
		carry = sum / 2
		r := strconv.Itoa(sum % 2)
		ret = r + ret

	}

	if carry == 1 {
		return "1" + ret
	}

	return ret
}

