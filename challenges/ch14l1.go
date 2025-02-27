package main

func getLast[T any](s []T) T {
	var myZero T
	if len(s) == 0 {
		return myZero
	}
	var last T = s[len(s)-1]
	return last
}
