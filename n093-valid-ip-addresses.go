package main
import (
    "strconv"
)

func restoreIpAddresses(s string) []string {
    res := []string{}
    
    var s1, s2, s3, s4 string
    for i := 1; i < 4 && i < len(s)-2; i++ {
        for j := i+1; j < i+4 && j < len(s)-1; j++ {
            for k := j+1; k < j+4 && k < len(s); k++ {
				s1 = s[0:i]
				s2 = s[i:j]
				s3 = s[j:k]
				s4 = s[k:]
                if isValid(s1) && isValid(s2) && isValid(s3) && isValid(s4) {
                    ip := s1 + "." + s2 + "." + s3 + "." + s4
                    res = append(res, ip)
                }
            }   
        }
    }
    
    return res
}
    
func isValid(s string) bool {
	if len(s) == 0 || len(s) > 3 || (s[0] == '0' && len(s) > 1)  {
		return false
	}
	
    i, err := strconv.Atoi(s)
	if err == nil && i <= 255 {
		return true
	}
	
	return false
}
