package main
import "sort"


func NonConstructibleChange(coins []int) int {
    // Sort the coins in ascending order.
    sort.Ints(coins)

    // Initialize the minimumChange to 1.
    minimumChange := 1

    // Iterate through the sorted coins.
    for _, coin := range coins {
        // If the current coin is less than or equal to minimumChange,
        // update minimumChange by adding the coin value.
        if coin <= minimumChange {
            minimumChange += coin
        } else {
            // If the current coin is greater than minimumChange,
            // it means there's a gap in the possible change values.
            // Return minimumChange as it cannot be created.
            return minimumChange
        }
    }

    return minimumChange
}