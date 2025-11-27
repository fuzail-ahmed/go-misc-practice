package main

func PrintLenCap(s []int) (int, int) {
	return len(s), cap(s)
}
