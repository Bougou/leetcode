package main

import (
    "fmt"
    "sort"
)

func groupAnagrams(strs []string) [][]string {
    strMap := map[string][]string{}
    
    for _, str := range strs {
        var intArr []int
        for _, c := range []byte(str) {
            intArr = append(intArr, int(c-0))
        }
        sort.Ints(intArr)
        fmt.Println(intArr)
        
        var byteArr []byte
        for _, c := range intArr {
            byteArr = append(byteArr, byte(c))
        }
        fmt.Println(byteArr)

        orderStr := string(byteArr)
        
        v, exists := strMap[orderStr]
        if !exists {
            strMap[orderStr] = []string{str}
        } else {
            strMap[orderStr] = append(v, str)
        }
        fmt.Println(strMap)
    }
    
    res := [][]string{}
    for _, v := range strMap {
        res = append(res, v)
    }
    
    return res
}