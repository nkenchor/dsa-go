package main
import(
    "sort"
)

func TwoNumberSum2Pointers(array []int, target int) []int {
    low:=0
    high:=len(array)-1
    sort.Ints(array)
    for low<high{
        currentSum:=array[low] + array[high] 
        if currentSum > target{
            high --
        }else if currentSum < target{
            low ++
        } else{
            return []int{array[low],array[high]}
        }
    }
    return []int{}
}
