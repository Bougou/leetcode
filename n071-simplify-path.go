package main

import "strings"

func simplifyPath(path string) string {
	// split string by '/'
	strList := strings.Split(path, "/")

	// remove '.' and space ''
	filterList := []string{}
	for _, v := range strList {
		if v != "." && v != "" {
			filterList = append(filterList, v)
		}
	}

	resStack := []string{}
	for _, v := range filterList {
		if v != ".." {
			resStack = append(resStack, v)
		} else {
			if len(resStack) > 0 {
				resStack = resStack[:len(resStack)-1]
			}
		}
	}

	return "/" + strings.Join(resStack, "/")
}
