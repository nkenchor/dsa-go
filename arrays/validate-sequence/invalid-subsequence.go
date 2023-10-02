package main
//time complexity: O(n)
func IsValidSubsequence(array []int, sequence []int) bool {
	// Write your code here.
    sequencePointer:=0

    for i:=0;i<len(array) && sequencePointer<len(sequence);i++{
        if array[i] == sequence[sequencePointer]{
            sequencePointer++
        }
    }
    return len(sequence)==sequencePointer
}
