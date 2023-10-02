package main

//time complexity O(n)
func SortedSquaredArray2Pointers(array []int) []int {
	startPointer:=0
    endPointer:=len(array) - 1
    sortedArray:= make([]int, len(array))

    for i:= len(array)-1; i>=0; i-- {
        endValue := array[endPointer] * array[endPointer]
        startValue := array[startPointer] * array[startPointer]

        if startValue > endValue {
            sortedArray[i] = startValue
            startPointer +=1
        } else {
            sortedArray[i] = endValue
            endPointer -=1
        }
    }
    
    return sortedArray
    
}
