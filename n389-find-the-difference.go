package main

func findTheDifference(s string, t string) byte {
	var ret byte

	for _, c := range append([]byte(s), []byte(t)...) {
		ret ^= c
	}

	return ret
}
