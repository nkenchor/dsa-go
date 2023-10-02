package main

import (
    "sort"
)
//time complexity 0(nlogn)
func SortedSquaredArraySort(array []int) []int {
	sortedSquare:= make([]int,len(array))

    for i:=0; i<len(array);i++ {
        sortedSquare[i]=array[i] * array[i]
    }

    sort.Ints(sortedSquare)
    return sortedSquare
}
