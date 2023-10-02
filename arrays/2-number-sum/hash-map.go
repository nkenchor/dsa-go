package main

func TwoNumberSumHashMap(array []int, target int) []int {
    // Create a hashmap to store numbers and their indices
    numMap := make(map[int]int)
    result := []int{}

    for i, num := range array {
        complement := target - num
        if index, found := numMap[complement]; found {
            // Found a pair, add them to the result
            result = append(result, array[index], num)
        }

        // Store the current number and its index in the hashmap
        numMap[num] = i
    }

    return result
}
